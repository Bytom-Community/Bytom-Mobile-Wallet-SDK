package api

import (
	"context"
	log "github.com/sirupsen/logrus"

	"github.com/bytom-community/mobile/sdk/account"
	"github.com/bytom-community/mobile/sdk/blockchain/query"
)

// POST /list-accounts
func (a *API) listAccounts(ctx context.Context, filter struct {
	ID string `json:"id"`
}) Response {
	accounts, err := a.Wallet.AccountMgr.ListAccounts(filter.ID)
	if err != nil {
		log.Errorf("listAccounts: %v", err)
		return NewErrorResponse(err)
	}

	annotatedAccounts := []query.AnnotatedAccount{}
	for _, acc := range accounts {
		annotatedAccounts = append(annotatedAccounts, *account.Annotated(acc))
	}

	return NewSuccessResponse(annotatedAccounts)
}
