package txbuilder

import (
	"github.com/bytom-community/mobile/sdk/protocol/vm"
	"github.com/bytom-community/mobile/sdk/protocol/vm/vmutil"
)

// Signature programs constrain how the signed inputs of a transaction
// in a template may be used, especially if the transaction is not yet
// complete.
//
// For example, suppose Alice wants to send Bob 80 EUR but only if Bob
// pays her 100 USD, and only if payment is made before next
// Tuesday. Alice constructs a partial transaction that includes her
// 80 EUR as one input, a payment to Bob as one output, _and_ a
// payment to Alice (of 100 USD) as one more output. She then
// constructs a program testing that the transaction includes all
// those components (and that the maxtime of the transaction is before
// next Tuesday) and signs a hash of that in order to unlock her 80
// EUR. She then passes the partial transaction template to Bob, who
// supplies his 100 USD input. Because of the signature program, Bob
// (or an eavesdropper) cannot use the signed 80-EUR input in any
// transaction other than one that pays 100 USD to Alice before
// Tuesday.
//
// This works because of Chain's convention for formatting of account
// control programs. The 80 EUR prevout being spent by Alice was paid
// to the program:
//   DUP TOALTSTACK SHA3 <pubkey1> <pubkey2> ... <pubkeyN> <quorum> <N> CHECKMULTISIG VERIFY FROMALTSTACK 0 CHECKPREDICATE
// which means that any attempt to spend it must be accompanied by a
// signed program that evaluates to true. The default program (for a
// complete transaction to which no other entries may be added) is
//   <txsighash> TXSIGHASH EQUAL
// which commits to the transaction as-is.

func buildSigProgram(tpl *Template, index uint32) ([]byte, error) {
	if !tpl.AllowAdditional {
		h := tpl.Hash(index)
		builder := vmutil.NewBuilder()
		builder.AddData(h.Bytes())
		builder.AddOp(vm.OP_TXSIGHASH).AddOp(vm.OP_EQUAL)

		return builder.Build()
	}
	constraints := make([]constraint, 0, 3+len(tpl.Transaction.Outputs))
	id := tpl.Transaction.Tx.InputIDs[index]
	if sp, err := tpl.Transaction.Tx.Spend(id); err == nil {
		constraints = append(constraints, outputIDConstraint(*sp.SpentOutputId))
	}

	for i, out := range tpl.Transaction.Outputs {
		c := &payConstraint{
			Index:       i,
			AssetAmount: out.AssetAmount,
			Program:     out.ControlProgram,
		}
		constraints = append(constraints, c)
	}
	var program []byte
	for i, c := range constraints {
		program = append(program, c.code()...)
		if i < len(constraints)-1 { // leave the final bool on top of the stack
			program = append(program, byte(vm.OP_VERIFY))
		}
	}
	return program, nil
}
