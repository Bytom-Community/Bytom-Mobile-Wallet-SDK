package txbuilder

// AddWitnessKeys adds a SignatureWitness with the given quorum and
// list of keys derived by applying the derivation path to each of the
// xpubs.

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
