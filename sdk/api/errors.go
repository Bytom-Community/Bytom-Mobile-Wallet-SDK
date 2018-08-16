package api

import (
	"github.com/bytom-community/mobile/sdk/account"
	"github.com/bytom-community/mobile/sdk/asset"
	"github.com/bytom-community/mobile/sdk/blockchain/pseudohsm"
	"github.com/bytom-community/mobile/sdk/blockchain/signers"
	"github.com/bytom-community/mobile/sdk/blockchain/txbuilder"
	"github.com/bytom-community/mobile/sdk/errors"
	"github.com/bytom-community/mobile/sdk/net/http/httperror"
)

var (
	// ErrDefault is default Bytom API Error
	ErrDefault = errors.New("Bytom API Error")
)

var respErrFormatter = map[error]httperror.Info{
	ErrDefault: {500, "BTM000", "Bytom API Error"},

	// Signers error namespace (2xx)
	signers.ErrBadQuorum: {400, "BTM200", "Quorum must be greater than 1 and less than or equal to the length of xpubs"},
	signers.ErrBadXPub:   {400, "BTM201", "Invalid xpub format"},
	signers.ErrNoXPubs:   {400, "BTM202", "At least one xpub is required"},
	signers.ErrBadType:   {400, "BTM203", "Retrieved type does not match expected type"},
	signers.ErrDupeXPub:  {400, "BTM204", "Root XPubs cannot contain the same key more than once"},

	// Transaction error namespace (7xx)
	// Build transaction error namespace (70x ~ 72x)
	txbuilder.ErrMissingFields: {400, "BTM707", "One or more fields are missing"},
	txbuilder.ErrBadAmount:     {400, "BTM708", "Invalid asset amount"},
	account.ErrFindAccount:     {400, "BTM709", "Not found account"},
	asset.ErrFindAsset:         {400, "BTM710", "Not found asset"},

	// Mock HSM error namespace (8xx)
	pseudohsm.ErrDuplicateKeyAlias:    {400, "BTM800", "Key Alias already exists"},
	pseudohsm.ErrInvalidAfter:         {400, "BTM801", "Invalid `after` in query"},
	pseudohsm.ErrLoadKey:              {400, "BTM802", "Key not found or wrong password"},
	pseudohsm.ErrTooManyAliasesToList: {400, "BTM803", "Requested key aliases exceeds limit"},
	pseudohsm.ErrDecrypt:              {400, "BTM804", "Could not decrypt key with given passphrase"},
}
