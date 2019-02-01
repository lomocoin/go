package resourceadapter

import (
	"context"

	"github.com/lomocoin/stellar-go/xdr"
	. "github.com/lomocoin/stellar-go/protocols/horizon"

)

func PopulateAsset(ctx context.Context, dest *Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
