package wallet

import (
	"encoding/json"
	"github.com/tendermint/tmlibs/db"
	aApi "github.com/bytom-community/mobile/sdk/api"
	"github.com/bytom-community/mobile/sdk/account"
	aWallet "github.com/bytom-community/mobile/sdk/wallet"
	"github.com/bytom-community/mobile/sdk/asset"
	"github.com/bytom-community/mobile/sdk/blockchain/pseudohsm"
	"github.com/bytom-community/mobile/sdk/blockchain/txbuilder"
	"github.com/bytom-community/mobile/sdk/crypto/ed25519/chainkd"
)

var api aApi.API

func InitWallet(storagePath string) {
	hsm := pseudohsm.New(storagePath)
	walletDB := db.NewDB("wallet", "leveldb", storagePath)
	accounts := account.NewManager(walletDB)
	assets := asset.NewRegistry(walletDB)
	wallet := aWallet.NewWallet(walletDB, accounts, assets, hsm)
	api = aApi.API{Wallet: wallet}
}

func CreateKey(alias string, password string) string {
	b, _ := json.Marshal(api.PseudohsmCreateKey(alias, password))
	return string(b)
}

func ListKey() string {
	b, _ := json.Marshal(api.PseudohsmListKeys())
	return string(b)
}

func CreateAccount(alias string, quorum int, rootXPub string) string {
	var XPubs []chainkd.XPub
	xpub := new(chainkd.XPub)
	xpub.UnmarshalText([]byte(rootXPub))
	XPubs = append(XPubs, *xpub)
	b, _ := json.Marshal(api.CreateAccount(XPubs, quorum, alias))
	return string(b)
}

func CreateAccountReceiver(accountID string, accountAlias string) string {
	b, _ := json.Marshal(api.CreateAccountReceiver(accountID, accountAlias))
	return string(b)
}

func ListAccounts() string {
	b, _ := json.Marshal(api.ListAccounts(""))
	return string(b)
}

func ListAddress(accountID string, accountAlias string) string {
	b, _ := json.Marshal(api.ListAddresses(accountID, accountAlias, 0, 0))
	return string(b)
}

func ResetKeyPassword(rootXPub string, oldPassword string, newPassword string) string {
	xpub := new(chainkd.XPub)
	xpub.UnmarshalText([]byte(rootXPub))
	b, _ := json.Marshal(api.PseudohsmResetPassword(*xpub, oldPassword, newPassword))
	return string(b)
}

func BackupWallet() string {
	b, _ := json.Marshal(api.BackupWalletImage())
	return string(b)
}

func RestoreWallet(walletImage string) string {
	var image aApi.WalletImage
	json.Unmarshal([]byte(walletImage), image)
	b, _ := json.Marshal(api.RestoreWalletImage(image))
	return string(b)
}

func SignTransaction(transaction string, password string) string {
	var tx txbuilder.Template
	json.Unmarshal([]byte(transaction), tx)
	b, _ := json.Marshal(api.PseudohsmSignTemplates(nil, password, tx))
	return string(b)
}
