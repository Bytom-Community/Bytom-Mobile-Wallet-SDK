package query

import (
	"encoding/json"

	"github.com/bytom-community/mobile/sdk/crypto/ed25519/chainkd"
	chainjson "github.com/bytom-community/mobile/sdk/encoding/json"
	"github.com/bytom-community/mobile/sdk/protocol/bc"
)

//AnnotatedTx means an annotated transaction.
type AnnotatedTx struct {
	ID                     bc.Hash            `json:"tx_id"`
	Timestamp              uint64             `json:"block_time"`
	BlockID                bc.Hash            `json:"block_hash"`
	BlockHeight            uint64             `json:"block_height"`
	Position               uint32             `json:"block_index"`
	BlockTransactionsCount uint32             `json:"block_transactions_count,omitempty"`
	Inputs                 []*AnnotatedInput  `json:"inputs"`
	Outputs                []*AnnotatedOutput `json:"outputs"`
	StatusFail             bool               `json:"status_fail"`
	Size                   uint64             `json:"size"`
}

//AnnotatedInput means an annotated transaction input.
type AnnotatedInput struct {
	Type            string             `json:"type"`
	AssetID         bc.AssetID         `json:"asset_id"`
	AssetAlias      string             `json:"asset_alias,omitempty"`
	AssetDefinition *json.RawMessage   `json:"asset_definition,omitempty"`
	Amount          uint64             `json:"amount"`
	IssuanceProgram chainjson.HexBytes `json:"issuance_program,omitempty"`
	ControlProgram  chainjson.HexBytes `json:"control_program,omitempty"`
	Address         string             `json:"address,omitempty"`
	SpentOutputID   *bc.Hash           `json:"spent_output_id,omitempty"`
	AccountID       string             `json:"account_id,omitempty"`
	AccountAlias    string             `json:"account_alias,omitempty"`
	Arbitrary       chainjson.HexBytes `json:"arbitrary,omitempty"`
	InputID         bc.Hash            `json:"input_id"`
}

//AnnotatedOutput means an annotated transaction output.
type AnnotatedOutput struct {
	Type            string             `json:"type"`
	OutputID        bc.Hash            `json:"id"`
	TransactionID   *bc.Hash           `json:"transaction_id,omitempty"`
	Position        int                `json:"position"`
	AssetID         bc.AssetID         `json:"asset_id"`
	AssetAlias      string             `json:"asset_alias,omitempty"`
	AssetDefinition *json.RawMessage   `json:"asset_definition,omitempty"`
	Amount          uint64             `json:"amount"`
	AccountID       string             `json:"account_id,omitempty"`
	AccountAlias    string             `json:"account_alias,omitempty"`
	ControlProgram  chainjson.HexBytes `json:"control_program"`
	Address         string             `json:"address,omitempty"`
}

//AnnotatedAccount means an annotated account.
type AnnotatedAccount struct {
	ID       string         `json:"id"`
	Alias    string         `json:"alias,omitempty"`
	XPubs    []chainkd.XPub `json:"xpubs"`
	Quorum   int            `json:"quorum"`
	KeyIndex uint64         `json:"key_index"`
}

//AnnotatedAsset means an annotated asset.
type AnnotatedAsset struct {
	ID              bc.AssetID         `json:"id"`
	Alias           string             `json:"alias,omitempty"`
	IssuanceProgram chainjson.HexBytes `json:"issuance_program"`
	Keys            []*AssetKey        `json:"keys"`
	Quorum          int                `json:"quorum"`
	Definition      *json.RawMessage   `json:"definition"`
}

//AssetKey means an asset key.
type AssetKey struct {
	RootXPub            chainkd.XPub         `json:"root_xpub"`
	AssetPubkey         chainjson.HexBytes   `json:"asset_pubkey"`
	AssetDerivationPath []chainjson.HexBytes `json:"asset_derivation_path"`
}

//AnnotatedUTXO means an annotated utxo.
type AnnotatedUTXO struct {
	Alias               string `json:"account_alias"`
	OutputID            string `json:"id"`
	AssetID             string `json:"asset_id"`
	AssetAlias          string `json:"asset_alias"`
	Amount              uint64 `json:"amount"`
	AccountID           string `json:"account_id"`
	Address             string `json:"address"`
	ControlProgramIndex uint64 `json:"control_program_index"`
	Program             string `json:"program"`
	SourceID            string `json:"source_id"`
	SourcePos           uint64 `json:"source_pos"`
	ValidHeight         uint64 `json:"valid_height"`
	Change              bool   `json:"change"`
}
