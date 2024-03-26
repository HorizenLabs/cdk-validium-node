package state

import (
	substrateTypes "github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// AttestationId struct
type AttestationId struct {
	BlockNumber   uint64
	AttestationId substrateTypes.U64
}
