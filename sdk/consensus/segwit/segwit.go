package segwit

import (
	"github.com/bytom-community/mobile/sdk/consensus"
	"github.com/bytom-community/mobile/sdk/protocol/vm"
)

func IsP2WPKHScript(prog []byte) bool {
	insts, err := vm.ParseProgram(prog)
	if err != nil {
		return false
	}
	if len(insts) != 2 {
		return false
	}
	if insts[0].Op > vm.OP_16 {
		return false
	}
	return insts[1].Op == vm.OP_DATA_20 && len(insts[1].Data) == consensus.PayToWitnessPubKeyHashDataSize
}

func IsP2WSHScript(prog []byte) bool {
	insts, err := vm.ParseProgram(prog)
	if err != nil {
		return false
	}
	if len(insts) != 2 {
		return false
	}
	if insts[0].Op > vm.OP_16 {
		return false
	}
	return insts[1].Op == vm.OP_DATA_32 && len(insts[1].Data) == consensus.PayToWitnessScriptHashDataSize
}
