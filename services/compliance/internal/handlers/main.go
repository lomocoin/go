package handlers

import (
	"strconv"
	"time"

	"github.com/lomocoin/stellar-go/clients/federation"
	"github.com/lomocoin/stellar-go/clients/stellartoml"
	"github.com/lomocoin/stellar-go/services/compliance/internal/config"
	"github.com/lomocoin/stellar-go/services/compliance/internal/crypto"
	"github.com/lomocoin/stellar-go/services/compliance/internal/db"
	"github.com/lomocoin/stellar-go/support/http"
)

// RequestHandler implements compliance server request handlers
type RequestHandler struct {
	Config                  *config.Config                 `inject:""`
	Client                  http.SimpleHTTPClientInterface `inject:""`
	Database                db.Database                    `inject:""`
	SignatureSignerVerifier crypto.SignerVerifierInterface `inject:""`
	StellarTomlResolver     stellartoml.ClientInterface    `inject:""`
	FederationResolver      federation.ClientInterface     `inject:""`
	NonceGenerator          NonceGeneratorInterface        `inject:""`
}

type NonceGeneratorInterface interface {
	Generate() string
}

type NonceGenerator struct{}

func (n *NonceGenerator) Generate() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

type TestNonceGenerator struct{}

func (n *TestNonceGenerator) Generate() string {
	return "nonce"
}
