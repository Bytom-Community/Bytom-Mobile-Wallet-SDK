package wallet

import (
	"github.com/bytom-community/mobile/sdk/blockchain/pseudohsm"
	a "github.com/bytom-community/mobile/sdk/api"
	"github.com/bytom-community/mobile/sdk/account"
	w "github.com/bytom-community/mobile/sdk/wallet"
	log "github.com/sirupsen/logrus"
	"github.com/tendermint/tmlibs/common"
	"github.com/tendermint/tmlibs/db"
	"encoding/json"
	"github.com/bytom-community/mobile/sdk/asset"
	"github.com/bytom-community/mobile/sdk/crypto/ed25519/chainkd"
)

func CreateAccount(keystorePath string, alias string, password string) string {

	hsm, err := pseudohsm.New(keystorePath)
	if err != nil {
		common.Exit(common.Fmt("initialize HSM failed: %v", err))
	}
	walletDB := db.NewDB("wallet", "leveldb", keystorePath)
	accounts := account.NewManager(walletDB)
	assets := asset.NewRegistry(walletDB)
	wallet, err := w.NewWallet(walletDB, accounts, assets, hsm)
	if err != nil {
		log.WithField("error", err).Error("init NewWallet")
	}
	api := &a.API{Wallet: wallet}
	xpub := [64]byte{}
	b, _ := json.Marshal(api.CreateAccount([]chainkd.XPub{xpub}, 1, alias))
	return string(b)
}
