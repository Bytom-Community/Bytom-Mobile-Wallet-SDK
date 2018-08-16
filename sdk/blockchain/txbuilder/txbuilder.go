// Package txbuilder builds a Chain Protocol transaction from
// a list of actions.
package txbuilder

import (
	"context"
	"github.com/bytom-community/mobile/sdk/errors"
)

// errors
var (
	//ErrBadTxInputIdx means unsigned tx input
	ErrBadTxInputIdx = errors.New("unsigned tx missing input")
	//ErrBadWitnessComponent means invalid witness component
	ErrBadWitnessComponent = errors.New("invalid witness component")
	//ErrBadAmount means invalid asset amount
	ErrBadAmount = errors.New("bad asset amount")
	//ErrMissingFields means missing required fields
	ErrMissingFields = errors.New("required field is missing")
)

// Sign will try to sign all the witness
func Sign(ctx context.Context, tpl *Template, auth string, signFn SignFunc) error {
	for i, sigInst := range tpl.SigningInstructions {
		for j, wc := range sigInst.WitnessComponents {
			switch sw := wc.(type) {
			case *SignatureWitness:
				err := sw.sign(ctx, tpl, uint32(i), auth, signFn)
				if err != nil {
					return errors.WithDetailf(err, "adding signature(s) to signature witness component %d of input %d", j, i)
				}
			case *RawTxSigWitness:
				err := sw.sign(ctx, tpl, uint32(i), auth, signFn)
				if err != nil {
					return errors.WithDetailf(err, "adding signature(s) to raw-signature witness component %d of input %d", j, i)
				}
			}
		}
	}
	return materializeWitnesses(tpl)
}
