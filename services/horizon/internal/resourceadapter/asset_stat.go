package resourceadapter

import (
	"context"

	"github.com/lomocoin/stellar-go/amount"
	. "github.com/lomocoin/stellar-go/protocols/horizon"
	"github.com/lomocoin/stellar-go/services/horizon/internal/db2/assets"
	"github.com/lomocoin/stellar-go/support/errors"
	"github.com/lomocoin/stellar-go/support/render/hal"
	"github.com/lomocoin/stellar-go/xdr"
)

// PopulateAssetStat fills out the details
//func PopulateAssetStat(
func PopulateAssetStat(
	ctx context.Context,
	res *AssetStat,
	row assets.AssetStatsR,
) (err error) {
	res.Asset.Type = row.Type
	res.Asset.Code = row.Code
	res.Asset.Issuer = row.Issuer
	res.Amount, err = amount.IntStringToAmount(row.Amount)
	if err != nil {
		return errors.Wrap(err, "Invalid amount in PopulateAssetStat")
	}
	res.NumAccounts = row.NumAccounts
	res.Flags = AccountFlags{
		(row.Flags & int8(xdr.AccountFlagsAuthRequiredFlag)) != 0,
		(row.Flags & int8(xdr.AccountFlagsAuthRevocableFlag)) != 0,
		(row.Flags & int8(xdr.AccountFlagsAuthImmutableFlag)) != 0,
	}
	res.PT = row.SortKey

	res.Links.Toml = hal.NewLink(row.Toml)
	return
}
