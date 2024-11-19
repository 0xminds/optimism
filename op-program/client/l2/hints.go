package l2

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	preimage "github.com/ethereum-optimism/optimism/op-preimage"
	"github.com/ethereum-optimism/optimism/op-service/eth"
)

const (
	HintL2BlockHeader    = "l2-block-header"
	HintL2Transactions   = "l2-transactions"
	HintL2Code           = "l2-code"
	HintL2StateNode      = "l2-state-node"
	HintL2Output         = "l2-output"
	HintL2AccountProof   = "l2-account-proof"
	HintL2PayloadWitness = "l2-payload-witness"
)

type BlockHeaderHint common.Hash

var _ preimage.Hint = BlockHeaderHint{}

func (l BlockHeaderHint) Hint() string {
	return HintL2BlockHeader + " " + (common.Hash)(l).String()
}

type TransactionsHint common.Hash

var _ preimage.Hint = TransactionsHint{}

func (l TransactionsHint) Hint() string {
	return HintL2Transactions + " " + (common.Hash)(l).String()
}

type CodeHint common.Hash

var _ preimage.Hint = CodeHint{}

func (l CodeHint) Hint() string {
	return HintL2Code + " " + (common.Hash)(l).String()
}

type StateNodeHint common.Hash

var _ preimage.Hint = StateNodeHint{}

func (l StateNodeHint) Hint() string {
	return HintL2StateNode + " " + (common.Hash)(l).String()
}

type L2OutputHint common.Hash

var _ preimage.Hint = L2OutputHint{}

func (l L2OutputHint) Hint() string {
	return HintL2Output + " " + (common.Hash)(l).String()
}

type AccountProofHint struct {
	BlockHash common.Hash
	Address   common.Address
}

var _ preimage.Hint = AccountProofHint{}

func (l AccountProofHint) Hint() string {
	hintData := append(l.BlockHash.Bytes(), l.Address.Bytes()...)

	return HintL2AccountProof + " " + hexutil.Encode(hintData)
}

type PayloadWitnessHint struct {
	ParentBlockHash   common.Hash           `json:"parentBlockHash"`
	PayloadAttributes eth.PayloadAttributes `json:"payloadAttributes"`
}

var _ preimage.Hint = PayloadWitnessHint{}

func (l PayloadWitnessHint) Hint() string {
	marshaled, err := json.Marshal(l)
	if err != nil {
		// should only happen if the struct is misconfigured
		panic(err)
	}

	return HintL2PayloadWitness + " " + hexutil.Encode(marshaled)
}
