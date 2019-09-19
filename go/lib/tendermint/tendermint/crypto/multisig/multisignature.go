package multisig

import (
	"github.com/terra-project/amino-js/go/lib/tendermint/tendermint/crypto/multisig/bitarray"
)

type Multisignature struct {
	BitArray *bitarray.CompactBitArray
	Sigs     [][]byte
}
