package resourceadapter

import (
	"context"

	"github.com/lomocoin/stellar-go/services/horizon/internal/db2/history"
	. "github.com/lomocoin/stellar-go/protocols/horizon"
)

func PopulateHistoryAccount(ctx context.Context, dest *HistoryAccount, row history.Account) {
	dest.ID = row.Address
	dest.AccountID = row.Address
}
