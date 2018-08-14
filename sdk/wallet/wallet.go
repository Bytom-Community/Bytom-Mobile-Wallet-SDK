package wallet

import (
	"github.com/tendermint/tmlibs/db"

	"github.com/bytom-community/mobile/sdk/account"
	"github.com/bytom-community/mobile/sdk/asset"
	"github.com/bytom-community/mobile/sdk/blockchain/pseudohsm"
)

//Wallet is related to storing account unspent outputs
type Wallet struct {
	DB         db.DB
	AccountMgr *account.Manager
	AssetReg   *asset.Registry
	Hsm        *pseudohsm.HSM
}

//NewWallet return a new wallet instance
func NewWallet(walletDB db.DB, account *account.Manager, asset *asset.Registry, hsm *pseudohsm.HSM) (*Wallet, error) {
	w := &Wallet{
		DB:         walletDB,
		AccountMgr: account,
		AssetReg:   asset,
		Hsm:        hsm,
	}
	return w, nil
}
