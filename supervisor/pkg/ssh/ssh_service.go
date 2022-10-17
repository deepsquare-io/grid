package ssh

import (
	"context"
	"encoding/base64"
	"io"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

const execTimeout = time.Duration(5 * time.Second)

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

func (s *Service) establish(user string) (session *ssh.Session, close func(), err error) {
	config := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
		Auth:            []ssh.AuthMethod{s.authMethod},
	}
	client, err := ssh.Dial("tcp", s.address, config)
	if err != nil {
		return nil, nil, err
	}
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

// ExecAs executes a command on the remote host with a timeout
func (s *Service) ExecAs(ctx context.Context, user string, cmd string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, execTimeout)
	defer cancel()
	stdChan := make(chan struct {
		string
		error
	})

	go func() {
		sess, close, err := s.establish(user)
		if err != nil {
			stdChan <- struct {
				string
				error
			}{
				"",
				err,
			}
			return
		}
		defer close()

		out, err := sess.CombinedOutput(cmd)
		stdChan <- struct {
			string
			error
		}{
			string(out),
			err,
		}
	}()

	select {
	case std := <-stdChan:
		if std.error != nil {
			logger.I.Error(
				"command failed",
				zap.Error(std.error),
				zap.Any("cmd", cmd),
				zap.String("output", std.string),
			)
			return std.string, std.error
		}
		return std.string, std.error
	case <-ctx.Done():
		logger.I.Error("command timed out",
			zap.Error(ctx.Err()),
			zap.Any("cmd", cmd),
		)
		return "", ctx.Err()
	}
}
