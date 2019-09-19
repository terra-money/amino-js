package slashing

import (
	sdk "github.com/terra-project/amino-js/go/lib/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgUnjail{}

type MsgUnjail struct {
	ValidatorAddr sdk.ValAddress `json:"address"` // address of the validator operator
}
