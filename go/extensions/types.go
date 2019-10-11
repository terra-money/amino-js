package extensions

import (
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
	"github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/x/auth"
)

var _ sdk.Msg = (*MsgSwap)(nil)
var _ sdk.Msg = (*MsgPricePrevote)(nil)
var _ sdk.Msg = (*MsgPriceVote)(nil)
var _ sdk.Msg = (*MsgDelegateFeederPermission)(nil)

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

type MsgPricePrevote struct {
	Hash      string         `json:"hash"` // hex string
	Denom     string         `json:"denom"`
	Feeder    sdk.AccAddress `json:"feeder"`
	Validator sdk.ValAddress `json:"validator"`
}

type MsgPriceVote struct {
	Price     sdk.Dec        `json:"price"` // the effective price of Luna in {Denom}
	Salt      string         `json:"salt"`
	Denom     string         `json:"denom"`
	Feeder    sdk.AccAddress `json:"feeder"`
	Validator sdk.ValAddress `json:"validator"`
}

type MsgDelegateFeederPermission struct {
	Operator     sdk.ValAddress `json:"operator"`
	FeedDelegate sdk.AccAddress `json:"feed_delegate"`
}

type PricePrevote struct {
	Hash        string         `json:"hash"`  // Vote hex hash to protect centralize data source problem
	Denom       string         `json:"denom"` // Ticker name of target fiat currency
	Voter       sdk.ValAddress `json:"voter"` // Voter val address
	SubmitBlock int64          `json:"submit_block"`
}

type PriceVote struct {
	Price sdk.Dec        `json:"price"` // Price of Luna in target fiat currency
	Denom string         `json:"denom"` // Ticker name of target fiat currency
	Voter sdk.ValAddress `json:"voter"` // voter val address of validator
}

type PriceBallot []PriceVote

const (
	TerraMsgSwap                      = "market/MsgSwap"
	TerraMsgPriceVote                 = "oracle/MsgPriceVote"
	TerraMsgPricePrevote              = "oracle/MsgPricePrevote"
	TerraMsgDelegateFeederPermission  = "oracle/MsgDelegateFeederPermission"
	TerraSchedule                     = "core/Schedule"
	TerraVestingSchedule              = "core/VestingSchedule"
	TerraBaseLazyGradedVestingAccount = "core/LazyGradedVestingAccount"
)
