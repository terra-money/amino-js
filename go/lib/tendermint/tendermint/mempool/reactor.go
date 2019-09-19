package mempool

import (
	"github.com/terra-project/amino-js/go/lib/tendermint/tendermint/types"
)

type PeerState interface {
}

type MempoolMessage interface{}

type TxMessage struct {
	Tx types.Tx
}
