package extensions

import (
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
	"github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/x/auth"
)

var _ sdk.Msg = (*MsgSwap)(nil)
var _ sdk.Msg = (*MsgExchangeRatePrevote)(nil)
var _ sdk.Msg = (*MsgExchangeRateVote)(nil)
var _ sdk.Msg = (*MsgDelegateFeedConsent)(nil)

type Schedule struct {
	StartTime int64   `json:"start_time"`
	EndTime   int64   `json:"end_time"`
	Ratio     sdk.Dec `json:"ratio"`
}

type VestingSchedule struct {
	Denom     string     `json:"denom"`
	Schedules []Schedule `json:"schedules"` // maps blocktime to percentage vested. Should sum to 1.
}

type BaseLazyGradedVestingAccount struct {
	*auth.BaseVestingAccount

	VestingSchedules []VestingSchedule `json:"vesting_schedules"`
}

type MsgSwap struct {
	Trader    sdk.AccAddress `json:"trader"`     // Address of the trader
	OfferCoin sdk.Coin       `json:"offer_coin"` // Coin being offered
	AskDenom  string         `json:"ask_denom"`  // Denom of the coin to swap to
}

type MsgExchangeRatePrevote struct {
	Hash      string         `json:"hash"` // hex string
	Denom     string         `json:"denom"`
	Feeder    sdk.AccAddress `json:"feeder"`
	Validator sdk.ValAddress `json:"validator"`
}

type MsgExchangeRateVote struct {
	ExchangeRate sdk.Dec        `json:"exchange_rate"` // the effective rate of Luna in {Denom}
	Salt         string         `json:"salt"`
	Denom        string         `json:"denom"`
	Feeder       sdk.AccAddress `json:"feeder"`
	Validator    sdk.ValAddress `json:"validator"`
}

type MsgDelegateFeedConsent struct {
	Operator  sdk.ValAddress `json:"operator"`
	Delegatee sdk.AccAddress `json:"delegatee"`
}

const (
	TerraMsgSwap                      = "market/MsgSwap"
	TerraMsgExchangeRateVote          = "oracle/MsgExchangeRateVote"
	TerraMsgExchangeRatePrevote       = "oracle/MsgExchangeRatePrevote"
	TerraMsgDelegateFeedConsent       = "oracle/MsgDelegateFeedConsent"
	TerraSchedule                     = "core/Schedule"
	TerraVestingSchedule              = "core/VestingSchedule"
	TerraBaseLazyGradedVestingAccount = "core/LazyGradedVestingAccount"
)
