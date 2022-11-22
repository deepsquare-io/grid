package ldap

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/deepsquare-io/the-grid/ldap-connector/config"
	"github.com/deepsquare-io/the-grid/ldap-connector/logger"
	"github.com/deepsquare-io/the-grid/ldap-connector/validate"
	"github.com/go-ldap/ldap/v3"
	"go.uber.org/zap"
)

type DataSource struct {
	url      string
	bindDN   string
	password string
	config   config.Config
	opts     []ldap.DialOpt
}

func New(
	url string,
	bindDN string,
	password string,
	config config.Config,
	insecure bool,
	caFile string,
) *DataSource {
	var tlsConfig tls.Config
	if !insecure {
		// Get the SystemCertPool, continue with an empty pool on error
		rootCAs, err := x509.SystemCertPool()
		if err != nil {
			logger.I.Error("failed to read system cert pool", zap.Error(err))
		}
		if rootCAs == nil {
			rootCAs = x509.NewCertPool()
		}
		if caFile != "" {
			certs, err := os.ReadFile(caFile)
			if err != nil {
				logger.I.Error("failed to read caFile", zap.String("caFile", caFile), zap.Error(err))
			}

			if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
				logger.I.Error("No certs appended, using system certs only")
			}
		}

		tlsConfig = tls.Config{
			RootCAs: rootCAs,
		}
	} else {
		tlsConfig = tls.Config{
			InsecureSkipVerify: true,
		}
	}
	return &DataSource{
		url:      url,
		bindDN:   bindDN,
		password: password,
		config:   config,
		opts: []ldap.DialOpt{
			ldap.DialWithTLSConfig(&tlsConfig),
		},
	}
}

func (d *DataSource) auth() (*ldap.Conn, error) {
	conn, err := ldap.DialURL(d.url, d.opts...)
	if err != nil {
		logger.I.Error("ldap dial failed", zap.Error(err))
		return nil, err
	}
	err = conn.Bind(d.bindDN, d.password)
	if err != nil {
		logger.I.Error("ldap auth failed", zap.Error(err))
		return nil, err
	}
	return conn, nil
}

func (d *DataSource) doWithTimeout(ctx context.Context, fn func() error) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	ch := make(chan error, 1)

	go func() {
		err := fn()
		ch <- err
	}()

	select {
	case <-ctx.Done():
		logger.I.Error("ldap context cancelled", zap.Error(ctx.Err()))
		return ctx.Err()
	case err := <-ch:
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *DataSource) HealthCheck(ctx context.Context) error {
	conn, err := d.auth()
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func (d *DataSource) CreateUser(ctx context.Context, user string) error {
	if errMsg := validate.LDAPUserIsValid(user); errMsg != "" {
		logger.I.Error("user is invalid", zap.Error(errors.New(errMsg)))
		return nil
	}

	conn, err := d.auth()
	if err != nil {
		return err
	}
	defer conn.Close()

	req := ldap.NewAddRequest(d.config.GetUserDN(user), []ldap.Control{})
	req.Attribute("objectClass", d.config.CreateUser.ObjectClasses)

	for _, attr := range d.config.CreateUser.UserNameAttributes {
		req.Attribute(attr, []string{user})
	}

	for key, value := range d.config.CreateUser.DefaultAttributes {
		req.Attribute(key, value)
	}

	if err := d.doWithTimeout(ctx, func() error {
		logger.I.Debug("ldap add", zap.Any("req", req))
		return conn.Add(req)
	}); err != nil {
		if strings.Contains(err.Error(), "LDAP Result Code 68") {
			logger.I.Info("user already exists, ignoring", zap.Error(err), zap.Any("req", req))
			return nil
		}
		logger.I.Error("ldap create user failed", zap.Error(err), zap.Any("req", req))
		return err
	}

	return nil
}

func (d *DataSource) AddUserToGroup(ctx context.Context, user string) error {
	if errMsg := validate.LDAPUserIsValid(user); errMsg != "" {
		logger.I.Error("user is invalid", zap.Error(errors.New(errMsg)))
		return nil
	}

	conn, err := d.auth()
	if err != nil {
		return err
	}
	defer conn.Close()

	req := ldap.NewModifyRequest(d.config.GroupDN, []ldap.Control{})
	for _, attr := range d.config.AddUserToGroup.MemberAttributes {
		req.Add(attr, []string{d.config.GetUserDN(user)})
	}

	if err := d.doWithTimeout(ctx, func() error {
		logger.I.Debug("ldap modify", zap.Any("req", req))
		return conn.Modify(req)
	}); err != nil {
		if strings.Contains(err.Error(), "LDAP Result Code 20") {
			logger.I.Info("user already exists, ignoring", zap.Error(err), zap.Any("req", req))
			return nil
		}
		logger.I.Error("ldap modify group failed", zap.Error(err), zap.Any("req", req))
		return err
	}

	return nil
}
