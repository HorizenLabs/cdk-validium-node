package types

import (
	"github.com/0xPolygonHermez/zkevm-node/aggregator/prover"
)

// FinalProofInputs struct
type FinalProofInputs struct {
	FinalProof       *prover.FinalProof
	NewLocalExitRoot []byte
	NewStateRoot     []byte
	AttestationId    uint64
	LeafCount        uint64
	LeafIndex        uint64
	MerklePath       []string
}
