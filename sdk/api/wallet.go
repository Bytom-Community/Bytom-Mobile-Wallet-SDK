package api

import (
	"github.com/bytom-community/mobile/sdk/account"
	"github.com/bytom-community/mobile/sdk/asset"
	"github.com/bytom-community/mobile/sdk/blockchain/pseudohsm"
	"github.com/bytom-community/mobile/sdk/errors"
)

// POST /wallet error
func (a *API) walletError() Response {
	return NewErrorResponse(errors.New("wallet not found, please check that the wallet is open"))
}

// WalletImage hold the ziped wallet data
type WalletImage struct {
	AccountImage *account.Image      `json:"account_image"`
	AssetImage   *asset.Image        `json:"asset_image"`
	KeyImages    *pseudohsm.KeyImage `json:"key_images"`
}

func (a *API) RestoreWalletImage(image WalletImage) Response {
	if err := a.Wallet.Hsm.Restore(image.KeyImages); err != nil {
		return NewErrorResponse(errors.Wrap(err, "restore key images"))
	}
	if err := a.Wallet.AssetReg.Restore(image.AssetImage); err != nil {
		return NewErrorResponse(errors.Wrap(err, "restore asset image"))
	}
	if err := a.Wallet.AccountMgr.Restore(image.AccountImage); err != nil {
		return NewErrorResponse(errors.Wrap(err, "restore account image"))
	}
	return NewSuccessResponse(nil)
}

func (a *API) BackupWalletImage() Response {
	keyImages, err := a.Wallet.Hsm.Backup()
	if err != nil {
		return NewErrorResponse(errors.Wrap(err, "backup key images"))
	}
	assetImage, err := a.Wallet.AssetReg.Backup()
	if err != nil {
		return NewErrorResponse(errors.Wrap(err, "backup asset image"))
	}
	accountImage, err := a.Wallet.AccountMgr.Backup()
	if err != nil {
		return NewErrorResponse(errors.Wrap(err, "backup account image"))
	}

	image := &WalletImage{
		KeyImages:    keyImages,
		AssetImage:   assetImage,
		AccountImage: accountImage,
	}
	return NewSuccessResponse(image)
}
