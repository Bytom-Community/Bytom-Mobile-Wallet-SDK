package txbuilder

import (
	"github.com/bytom-community/mobile/sdk/errors"
)

var (
	// ErrMissingRawTx means missing transaction
	ErrMissingRawTx = errors.New("missing raw tx")
	// ErrBadInstructionCount means too many signing instructions compare with inputs
	ErrBadInstructionCount = errors.New("too many signing instructions in template")
)
