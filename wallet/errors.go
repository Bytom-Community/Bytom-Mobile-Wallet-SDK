package wallet

import (
	"github.com/bytom-community/mobile/wallet/pseudohsm"
	"github.com/bytom-community/mobile/wallet/errors"
	"github.com/bytom-community/mobile/wallet/net/http/httperror"
)

var (
	// ErrDefault is default Bytom API Error
	ErrDefault = errors.New("Bytom API Error")
)

var respErrFormatter = map[error]httperror.Info{
	ErrDefault: {500, "BTM000", "Bytom API Error"},

	// Mock HSM error namespace (8xx)
	pseudohsm.ErrDuplicateKeyAlias:    {400, "BTM800", "Key Alias already exists"},
	pseudohsm.ErrInvalidAfter:         {400, "BTM801", "Invalid `after` in query"},
	pseudohsm.ErrLoadKey:              {400, "BTM802", "Key not found or wrong password"},
	pseudohsm.ErrTooManyAliasesToList: {400, "BTM803", "Requested key aliases exceeds limit"},
	pseudohsm.ErrDecrypt:              {400, "BTM804", "Could not decrypt key with given passphrase"},
}
