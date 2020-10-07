module github.com/cosmos/amino-js/go/js

go 1.12

require (
	github.com/cosmos/amino-js/go/src v0.0.0
	github.com/gopherjs/gopherjs v0.0.0-20190812055313-0a3de5b74c85
	github.com/tendermint/go-amino v0.15.1
)

replace github.com/cosmos/amino-js/go/extensions => ../extensions

replace github.com/cosmos/amino-js/go/src => ../src

replace github.com/cosmos/amino-js/go/lib => ../lib
