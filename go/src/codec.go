package src

import (
	"github.com/terra-project/amino-js/go/extensions"
	"github.com/terra-project/amino-js/go/lib"
	"github.com/tendermint/go-amino"
)

var codec *amino.Codec

func init() {
	codec = amino.NewCodec()
	lib.RegisterCodec(codec)
	extensions.RegisterCodec(codec)
	codec.Seal()
}
