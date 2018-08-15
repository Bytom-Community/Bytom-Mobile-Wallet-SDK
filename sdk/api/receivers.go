package api

import (
	"github.com/bytom-community/mobile/sdk/blockchain/txbuilder"
)

func (a *API) CreateAccountReceiver(accountID string, accountAlias string) Response {
	if accountAlias != "" {
		account, err := a.Wallet.AccountMgr.FindByAlias(accountAlias)
		if err != nil {
			return NewErrorResponse(err)
		}

		accountID = account.ID
	}

	program, err := a.Wallet.AccountMgr.CreateAddress(accountID, false)
	if err != nil {
		return NewErrorResponse(err)
	}

	return NewSuccessResponse(&txbuilder.Receiver{
		ControlProgram: program.ControlProgram,
		Address:        program.Address,
	})
}
