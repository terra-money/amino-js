package evidence

import (
	"github.com/terra-project/amino-js/go/lib/tendermint/tendermint/types"
)

type PeerState interface {
}

type EvidenceMessage interface {
}

type EvidenceListMessage struct {
	Evidence []types.Evidence
}
