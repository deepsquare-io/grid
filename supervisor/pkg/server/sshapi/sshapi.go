package sshapi

import (
	"context"
	"encoding/base64"

	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/supervisor/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

type Server struct {
	supervisorv1alpha1.UnimplementedSshAPIServer
	pub ssh.PublicKey
}

func New(b64pk string) *Server {
	pk, err := base64.StdEncoding.DecodeString(b64pk)
	if err != nil {
		logger.I.Panic("failed to decode key", zap.Error(err))
	}

	signer, err := ssh.ParsePrivateKey(pk)
	if err != nil {
		logger.I.Panic("couldn't parse private key", zap.Error(err))
	}

	return &Server{
		pub: signer.PublicKey(),
	}
}

func (s *Server) FetchAuthorizedKeys(
	ctx context.Context,
	req *supervisorv1alpha1.FetchAuthorizedKeysRequest,
) (*supervisorv1alpha1.FetchAuthorizedKeysResponse, error) {
	return &supervisorv1alpha1.FetchAuthorizedKeysResponse{
		AuthorizedKeys: string(ssh.MarshalAuthorizedKey(s.pub)),
	}, nil
}
