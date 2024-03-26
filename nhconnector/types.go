package nhconnector

import "github.com/centrifuge/go-substrate-rpc-client/v4/types"

// NH Event for Proof verified
type PoeNewElement struct {
	Phase         types.Phase
	Value         types.Hash
	AttestationId types.U64
	Topics        []types.Hash
}

// NH Event for new Root (Attestation) created
type PoeNewAttestation struct {
	Phase       types.Phase
	Id          types.U64
	Attestation types.Hash
	Topics      []types.Hash
}

// NH Events map
type NHEventsRecords struct {
	Poe_NewElement     []PoeNewElement     `test-gen-blockchain:"nh-core"`
	Poe_NewAttestation []PoeNewAttestation `test-gen-blockchain:"nh-core"`
}

// NH RPC Requests
const (
	PoeProofPathRequest = "poe_proofPath"
)

// NH RPC Responses
type PoeProofPathRequestResponse struct {
	Root           string    `json:"root"`
	LeafIndex      types.U64 `json:"leaf_index"`
	Leaf           string    `json:"leaf"`
	Proof          []string  `json:"proof"`
	NumberOfLeaves types.U64 `json:"number_of_leaves"`
}

// NH Extrinsic endpoints
const (
	SubmitProofExtrinsic = "SettlementFFlonkPallet.submit_proof"
)
