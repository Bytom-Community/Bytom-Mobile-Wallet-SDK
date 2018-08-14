package txbuilder

import (
	"encoding/json"

	"github.com/bytom-community/mobile/sdk/crypto/ed25519/chainkd"
	chainjson "github.com/bytom-community/mobile/sdk/encoding/json"
	"github.com/bytom-community/mobile/sdk/errors"
)

// AddWitnessKeys adds a SignatureWitness with the given quorum and
// list of keys derived by applying the derivation path to each of the
// xpubs.
func (si *SigningInstruction) AddWitnessKeys(xpubs []chainkd.XPub, path [][]byte, quorum int) {
	hexPath := make([]chainjson.HexBytes, 0, len(path))
	for _, p := range path {
		hexPath = append(hexPath, p)
	}

	keyIDs := make([]keyID, 0, len(xpubs))
	for _, xpub := range xpubs {
		keyIDs = append(keyIDs, keyID{xpub, hexPath})
	}

	sw := &SignatureWitness{
		Quorum: quorum,
		Keys:   keyIDs,
	}
	si.WitnessComponents = append(si.WitnessComponents, sw)
}

// AddRawWitnessKeys adds a SignatureWitness with the given quorum and
// list of keys derived by applying the derivation path to each of the
// xpubs.
func (si *SigningInstruction) AddRawWitnessKeys(xpubs []chainkd.XPub, path [][]byte, quorum int) {
	hexPath := make([]chainjson.HexBytes, 0, len(path))
	for _, p := range path {
		hexPath = append(hexPath, p)
	}

	keyIDs := make([]keyID, 0, len(xpubs))
	for _, xpub := range xpubs {
		keyIDs = append(keyIDs, keyID{xpub, hexPath})
	}

	sw := &RawTxSigWitness{
		Quorum: quorum,
		Keys:   keyIDs,
	}
	si.WitnessComponents = append(si.WitnessComponents, sw)
}

// SigningInstruction gives directions for signing inputs in a TxTemplate.
type SigningInstruction struct {
	Position          uint32             `json:"position"`
	WitnessComponents []witnessComponent `json:"witness_components,omitempty"`
}

// witnessComponent is the abstract type for the parts of a
// SigningInstruction.  Each witnessComponent produces one or more
// arguments for a VM program via its materialize method. Concrete
// witnessComponent types include SignatureWitness and dataWitness.
type witnessComponent interface {
	materialize(*[][]byte) error
}

// UnmarshalJSON unmarshal SigningInstruction
func (si *SigningInstruction) UnmarshalJSON(b []byte) error {
	var pre struct {
		Position          uint32            `json:"position"`
		WitnessComponents []json.RawMessage `json:"witness_components"`
	}
	err := json.Unmarshal(b, &pre)
	if err != nil {
		return err
	}

	si.Position = pre.Position
	for i, wc := range pre.WitnessComponents {
		var t struct {
			Type string
		}
		err = json.Unmarshal(wc, &t)
		if err != nil {
			return errors.Wrapf(err, "unmarshaling error on witness component %d, input %s", i, wc)
		}
		switch t.Type {
		case "data":
			var d struct {
				Value chainjson.HexBytes
			}
			err = json.Unmarshal(wc, &d)
			if err != nil {
				return errors.Wrapf(err, "unmarshaling error on witness component %d, type data, input %s", i, wc)
			}
			si.WitnessComponents = append(si.WitnessComponents, DataWitness(d.Value))

		case "signature":
			var s SignatureWitness
			err = json.Unmarshal(wc, &s)
			if err != nil {
				return errors.Wrapf(err, "unmarshaling error on witness component %d, type signature, input %s", i, wc)
			}
			si.WitnessComponents = append(si.WitnessComponents, &s)

		case "raw_tx_signature":
			var s RawTxSigWitness
			err = json.Unmarshal(wc, &s)
			if err != nil {
				return errors.Wrapf(err, "unmarshaling error on witness component %d, type raw_signature, input %s", i, wc)
			}
			si.WitnessComponents = append(si.WitnessComponents, &s)

		default:
			return errors.WithDetailf(ErrBadWitnessComponent, "witness component %d has unknown type '%s'", i, t.Type)
		}
	}
	return nil
}
