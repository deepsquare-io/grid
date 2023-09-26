package ssh

import (
	"context"
	"encoding/base64"
	"io"
	"net"
	"time"

	"github.com/deepsquare-io/grid/supervisor/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

type Service struct {
	address    string
	authMethod ssh.AuthMethod
}

func New(
	address string,
	pkB64 string,
) *Service {
	pk, err := base64.StdEncoding.DecodeString(pkB64)
	if err != nil {
		logger.I.Panic("failed to decode key", zap.Error(err))
	}

	signer, err := ssh.ParsePrivateKey(pk)
	if err != nil {
		logger.I.Panic("couldn't parse private key", zap.Error(err))
	}

	return &Service{
		address:    address,
		authMethod: ssh.PublicKeys(signer),
	}
}

func (s *Service) establish(
	ctx context.Context,
	user string,
) (session *ssh.Session, close func(), err error) {
	config := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
		Auth:            []ssh.AuthMethod{s.authMethod},
	}
	d := net.Dialer{Timeout: config.Timeout}
	conn, err := d.DialContext(ctx, "tcp", s.address)
	if err != nil {
		return nil, nil, err
	}
	c, chans, reqs, err := ssh.NewClientConn(conn, s.address, config)
	if err != nil {
		return nil, nil, err
	}
	client := ssh.NewClient(c, chans, reqs)
	session, err = client.NewSession()
	if err != nil {
		if err := client.Close(); err != nil && err != io.EOF {
			logger.I.Warn("closing SSH client thrown an error", zap.Error(err))
		}
		return nil, nil, err
	}

	return session, func() {
		if err := session.Close(); err != nil && err != io.EOF {
			logger.I.Warn("closing SSH session thrown an error", zap.Error(err))
		}
		if err := client.Close(); err != nil && err != io.EOF {
			logger.I.Warn("closing SSH client thrown an error", zap.Error(err))
		}
	}, nil
}

// ExecAs executes a command on the remote host
func (s *Service) ExecAs(ctx context.Context, user string, cmd string) (string, error) {
	sess, close, err := s.establish(ctx, user)
	if err != nil {
		return "", err
	}
	defer close()

	logger.I.Debug(
		"called ssh command",
		zap.String("cmd", cmd),
		zap.String("user", user),
	)
	out, err := sess.CombinedOutput(cmd)
	return string(out), err
}
