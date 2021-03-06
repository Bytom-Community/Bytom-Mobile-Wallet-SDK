package api

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/bytom-community/mobile/sdk/blockchain/txbuilder"
	"github.com/bytom-community/mobile/sdk/crypto/ed25519/chainkd"
)

func (a *API) PseudohsmCreateKey(alias string, password string) Response {
	xpub, err := a.Wallet.Hsm.XCreate(alias, password)
	if err != nil {
		return NewErrorResponse(err)
	}
	return NewSuccessResponse(xpub)
}

func (a *API) PseudohsmListKeys() Response {
	return NewSuccessResponse(a.Wallet.Hsm.ListKeys())
}

func (a *API) pseudohsmDeleteKey(ctx context.Context, x struct {
	Password string       `json:"password"`
	XPub     chainkd.XPub `json:"xpub"`
}) Response {
	if err := a.Wallet.Hsm.XDelete(x.XPub, x.Password); err != nil {
		return NewErrorResponse(err)
	}
	return NewSuccessResponse(nil)
}

type signResp struct {
	Tx           *txbuilder.Template `json:"transaction"`
	SignComplete bool                `json:"sign_complete"`
}

func (a *API) PseudohsmSignTemplates(ctx context.Context, password string, txs txbuilder.Template) Response {
	if err := txbuilder.Sign(ctx, &txs, password, a.pseudohsmSignTemplate); err != nil {
		log.WithField("build err", err).Error("fail on sign transaction.")
		return NewErrorResponse(err)
	}
	log.Info("Sign Transaction complete.")
	return NewSuccessResponse(&signResp{Tx: &txs, SignComplete: txbuilder.SignProgress(&txs)})
}

func (a *API) pseudohsmSignTemplate(ctx context.Context, xpub chainkd.XPub, path [][]byte, data [32]byte, password string) ([]byte, error) {
	return a.Wallet.Hsm.XSign(xpub, path, data[:], password)
}

// ResetPasswordResp is response for reset key password
type ResetPasswordResp struct {
	Changed bool `json:"changed"`
}

func (a *API) PseudohsmResetPassword(xPub chainkd.XPub, oldPassword string, newPassword string) Response {
	resp := &ResetPasswordResp{Changed: false}
	if err := a.Wallet.Hsm.ResetPassword(xPub, oldPassword, newPassword); err != nil {
		return NewSuccessResponse(resp)
	}
	resp.Changed = true
	return NewSuccessResponse(resp)
}

// CheckPasswordResp is response for check key password
type CheckPasswordResp struct {
	CheckResult bool `json:"check_result"`
}

func (a *API) pseudohsmCheckPassword(ctx context.Context, ins struct {
	XPub     chainkd.XPub `json:"xpub"`
	Password string       `json:"password"`
}) Response {
	resp := &CheckPasswordResp{CheckResult: false}
	if _, err := a.Wallet.Hsm.LoadChainKDKey(ins.XPub, ins.Password); err != nil {
		return NewSuccessResponse(resp)
	}
	resp.CheckResult = true
	return NewSuccessResponse(resp)
}
