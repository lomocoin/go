package handlers

import (
	"github.com/lomocoin/stellar-go/clients/federation"
	"github.com/lomocoin/stellar-go/clients/horizon"
	"github.com/lomocoin/stellar-go/clients/stellartoml"
	"github.com/lomocoin/stellar-go/services/bridge/internal/config"
	"github.com/lomocoin/stellar-go/services/bridge/internal/db"
	"github.com/lomocoin/stellar-go/services/bridge/internal/listener"
	"github.com/lomocoin/stellar-go/services/bridge/internal/submitter"
	"github.com/lomocoin/stellar-go/support/http"
)

// RequestHandler implements bridge server request handlers
type RequestHandler struct {
	Config               *config.Config                          `inject:""`
	Client               http.SimpleHTTPClientInterface          `inject:""`
	Horizon              horizon.ClientInterface                 `inject:""`
	Database             db.Database                             `inject:""`
	StellarTomlResolver  stellartoml.ClientInterface             `inject:""`
	FederationResolver   federation.ClientInterface              `inject:""`
	TransactionSubmitter submitter.TransactionSubmitterInterface `inject:""`
	PaymentListener      *listener.PaymentListener               `inject:""`
}

func (rh *RequestHandler) isAssetAllowed(code string, issuer string) bool {
	for _, asset := range rh.Config.Assets {
		if asset.Code == code && asset.Issuer == issuer {
			return true
		}
	}
	return false
}
