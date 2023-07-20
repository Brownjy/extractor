package grafana

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/specs-actors/v3/actors/builtin"
	"time"
)

var genesis *types.BlockHeader

func time2epoch(t time.Time) abi.ChainEpoch {
	return abi.ChainEpoch((uint64(t.Unix()) - genesis.Timestamp) / builtin.EpochDurationSeconds)
}
func epoch2time(h abi.ChainEpoch) time.Time {
	return time.Unix(int64(genesis.Timestamp)+int64(h)*builtin.EpochDurationSeconds, 0)
}
