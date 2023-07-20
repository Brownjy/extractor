package segment

import (
	"extractor/common"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"sync"
)

// Anchor contains most significant info about a tipset
type Anchor struct {
	Epoch abi.ChainEpoch
	TSK   types.TipSetKey
	State cid.Cid
}
type Options struct {
	name string
	rdb  common.DocumentDB

	bound struct {
		sync.RWMutex
		Boundary
	}
}

func DefaultSegment() Options {
	return Options{
		name: "test-segment",
		bound: struct {
			sync.RWMutex
			Boundary
		}{},
	}
}
