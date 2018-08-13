package wallet

import (
	"github.com/bytom-community/mobile/wallet/pseudohsm"
	"encoding/json"
)

func CreateWallet(keystorePath string, alias string, password string) string {
	hsm, _ := pseudohsm.New(keystorePath)
	api := &API{Hsm: hsm}
	b, _ := json.Marshal(api.PseudohsmCreateKey(alias, password))
	return string(b)
}