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

//go:build unit

package sshapi_test

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"testing"

	supervisorv1alpha1 "github.com/deepsquare-io/grid/supervisor/generated/supervisor/v1alpha1"
	"github.com/deepsquare-io/grid/supervisor/pkg/server/sshapi"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/ssh"
)

type ServerTestSuite struct {
	suite.Suite
	impl *sshapi.Server
	pub  ed25519.PublicKey
	pk   ed25519.PrivateKey
}

func (suite *ServerTestSuite) BeforeTest(suiteName, testName string) {
	pub, pk, err := ed25519.GenerateKey(nil)
	suite.pub = pub
	suite.pk = pk

	// Create a PEM block for the private key
	b, err := x509.MarshalPKCS8PrivateKey(pk)
	suite.Require().NoError(err)

	privateKeyPEM := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: b,
	}

	// Serialize the PEM block
	privateKeyPEMBytes := pem.EncodeToMemory(privateKeyPEM)

	// Base64 encode the serialized PEM block
	privateKeyBase64 := base64.StdEncoding.EncodeToString(privateKeyPEMBytes)

	suite.Require().NoError(err)
	suite.impl = sshapi.New(privateKeyBase64)
}

func (suite *ServerTestSuite) TestFetchAuthorizedKeys() {
	ctx := context.Background()
	// Arrange
	pub, err := ssh.NewPublicKey(suite.pub)
	suite.Require().NoError(err)

	// Act
	resp, err := suite.impl.FetchAuthorizedKeys(
		ctx,
		&supervisorv1alpha1.FetchAuthorizedKeysRequest{},
	)
	suite.Require().NoError(err)
	suite.Require().Equal(&supervisorv1alpha1.FetchAuthorizedKeysResponse{
		AuthorizedKeys: string(ssh.MarshalAuthorizedKey(pub)),
	}, resp)
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, &ServerTestSuite{})
}
