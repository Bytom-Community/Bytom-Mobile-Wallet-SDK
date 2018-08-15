package api

import (
	"context"
	"encoding/hex"
	"sort"

	log "github.com/sirupsen/logrus"

	"github.com/bytom-community/mobile/sdk/account"
	"github.com/bytom-community/mobile/sdk/common"
	"github.com/bytom-community/mobile/sdk/consensus"
	"github.com/bytom-community/mobile/sdk/crypto/ed25519/chainkd"
	"github.com/bytom-community/mobile/sdk/protocol/vm/vmutil"
)

// POST /create-account
func (a *API) CreateAccount(RootXPubs []chainkd.XPub, Quorum int, Alias string) Response {
	acc, err := a.Wallet.AccountMgr.Create(RootXPubs, Quorum, Alias)
	if err != nil {
		return NewErrorResponse(err)
	}

	annotatedAccount := account.Annotated(acc)
	log.WithField("account ID", annotatedAccount.ID).Info("Created account")

	return NewSuccessResponse(annotatedAccount)
}

// AccountInfo is request struct for deleteAccount
type AccountInfo struct {
	Info string `json:"account_info"`
}

// POST /delete-account
func (a *API) deleteAccount(ctx context.Context, in AccountInfo) Response {
	if err := a.Wallet.AccountMgr.DeleteAccount(in.Info); err != nil {
		return NewErrorResponse(err)
	}
	return NewSuccessResponse(nil)
}

type validateAddressResp struct {
	Valid   bool `json:"valid"`
	IsLocal bool `json:"is_local"`
}

// POST /validate-address
func (a *API) validateAddress(ctx context.Context, ins struct {
	Address string `json:"address"`
}) Response {
	resp := &validateAddressResp{
		Valid:   false,
		IsLocal: false,
	}
	address, err := common.DecodeAddress(ins.Address, &consensus.ActiveNetParams)
	if err != nil {
		return NewSuccessResponse(resp)
	}

	redeemContract := address.ScriptAddress()
	program := []byte{}
	switch address.(type) {
	case *common.AddressWitnessPubKeyHash:
		program, err = vmutil.P2WPKHProgram(redeemContract)
	case *common.AddressWitnessScriptHash:
		program, err = vmutil.P2WSHProgram(redeemContract)
	default:
		return NewSuccessResponse(resp)
	}
	if err != nil {
		return NewSuccessResponse(resp)
	}

	resp.Valid = true
	resp.IsLocal = a.Wallet.AccountMgr.IsLocalControlProgram(program)
	return NewSuccessResponse(resp)
}

type addressResp struct {
	AccountAlias   string `json:"account_alias"`
	AccountID      string `json:"account_id"`
	Address        string `json:"address"`
	ControlProgram string `json:"control_program"`
	Change         bool   `json:"change"`
	KeyIndex       uint64 `json:"-"`
}

// SortByIndex implements sort.Interface for addressResp slices
type SortByIndex []addressResp

func (a SortByIndex) Len() int           { return len(a) }
func (a SortByIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByIndex) Less(i, j int) bool { return a[i].KeyIndex < a[j].KeyIndex }

func (a *API) ListAddresses(accountID string, accountAlias string, from uint, count uint, ) Response {
	var target *account.Account
	if accountAlias != "" {
		acc, err := a.Wallet.AccountMgr.FindByAlias(accountAlias)
		if err != nil {
			return NewErrorResponse(err)
		}
		target = acc
	} else {
		acc, err := a.Wallet.AccountMgr.FindByID(accountID)
		if err != nil {
			return NewErrorResponse(err)
		}
		target = acc
	}

	cps, err := a.Wallet.AccountMgr.ListControlProgram()
	if err != nil {
		return NewErrorResponse(err)
	}

	addresses := []addressResp{}
	for _, cp := range cps {
		if cp.Address == "" || cp.AccountID != target.ID {
			continue
		}
		addresses = append(addresses, addressResp{
			AccountAlias:   target.Alias,
			AccountID:      cp.AccountID,
			Address:        cp.Address,
			ControlProgram: hex.EncodeToString(cp.ControlProgram),
			Change:         cp.Change,
			KeyIndex:       cp.KeyIndex,
		})
	}

	// sort AddressResp by KeyIndex
	sort.Sort(SortByIndex(addresses))
	start, end := getPageRange(len(addresses), from, count)
	return NewSuccessResponse(addresses[start:end])
}
