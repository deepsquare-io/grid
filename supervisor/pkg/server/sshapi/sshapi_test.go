//go:build unit

package sshapi_test

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"testing"

	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/generated/supervisor/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server/sshapi"
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
