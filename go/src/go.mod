module github.com/cosmos/amino-js/go/src

go 1.13

require (
	github.com/cosmos/amino-js/go/extensions v0.0.0
	github.com/cosmos/amino-js/go/lib v0.0.0
	github.com/google/gofuzz v1.0.0 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/tendermint/go-amino v0.15.1
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace github.com/cosmos/amino-js/go/extensions => ../extensions

replace github.com/cosmos/amino-js/go/lib => ../lib
