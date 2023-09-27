// Copyright (C) 2023 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package sshapi

import (
	"context"
	"encoding/base64"

	supervisorv1alpha1 "github.com/deepsquare-io/grid/supervisor/generated/supervisor/v1alpha1"
	"github.com/deepsquare-io/grid/supervisor/logger"
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
