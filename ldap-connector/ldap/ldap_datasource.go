package ldap

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"

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
		certs, err := os.ReadFile(caFile)
		if err != nil {
			logger.I.Error("failed to read caFile", zap.String("caFile", caFile), zap.Error(err))
		}

		if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
			logger.I.Error("No certs appended, using system certs only")
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
		url:    url,
		config: config,
		opts: []ldap.DialOpt{
			ldap.DialWithTLSConfig(&tlsConfig),
		},
	}
}

func (d *DataSource) auth() (*ldap.Conn, error) {
	conn, err := ldap.DialURL(d.url, d.opts...)
	if err != nil {
		return nil, err
	}
	err = conn.Bind(d.bindDN, d.password)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (d *DataSource) CreateUser(ctx context.Context, user string) error {
	if errMsg := validate.LDAPUserIsValid(user); errMsg != "" {
		return errors.New(errMsg)
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

	return conn.Add(req)
}

func (d *DataSource) AddUserToGroup(ctx context.Context, user string) error {
	if errMsg := validate.LDAPUserIsValid(user); errMsg != "" {
		return errors.New(errMsg)
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

	return conn.Modify(req)
}
