package txbuilder

import (
	chainjson "github.com/bytom-community/mobile/sdk/encoding/json"
	"github.com/bytom-community/mobile/sdk/protocol/bc"
	"github.com/bytom-community/mobile/sdk/protocol/bc/types"
)

// Template represents a partially- or fully-signed transaction.
type Template struct {
	Transaction         *types.Tx             `json:"raw_transaction"`
	SigningInstructions []*SigningInstruction `json:"signing_instructions"`

	// AllowAdditional affects whether Sign commits to the tx sighash or
	// to individual details of the tx so far. When true, signatures
	// commit to tx details, and new details may be added but existing
	// ones cannot be changed. When false, signatures commit to the tx
	// as a whole, and any change to the tx invalidates the signature.
	AllowAdditional bool `json:"allow_additional_actions"`
}

// Hash return sign hash
func (t *Template) Hash(idx uint32) bc.Hash {
	return t.Transaction.SigHash(idx)
}

// Receiver encapsulates information about where to send assets.
type Receiver struct {
	ControlProgram chainjson.HexBytes `json:"control_program,omitempty"`
	Address        string             `json:"address,omitempty"`
}
