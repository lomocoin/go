package resourceadapter

import (
	"github.com/lomocoin/stellar-go/services/horizon/internal/db2/core"
	. "github.com/lomocoin/stellar-go/protocols/horizon"
)

func PopulateAccountFlags(dest *AccountFlags, row core.Account) {
	dest.AuthRequired = row.IsAuthRequired()
	dest.AuthRevocable = row.IsAuthRevocable()
	dest.AuthImmutable = row.IsAuthImmutable()
}
