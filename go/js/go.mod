module github.com/cosmos/amino-js/go/js

go 1.12

require (
	github.com/cosmos/amino-js/go/src v0.0.0
	github.com/gopherjs/gopherjs v0.0.0-20190915194858-d3ddacdb130f
	github.com/tendermint/go-amino v0.15.0
)

replace github.com/cosmos/amino-js/go/extensions => ../extensions

replace github.com/cosmos/amino-js/go/src => ../src

replace github.com/cosmos/amino-js/go/lib => ../lib
