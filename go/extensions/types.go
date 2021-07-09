package extensions

import (
	"time"

	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
	"github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/x/auth"
)

var _ sdk.Msg = (*MsgSwap)(nil)
var _ sdk.Msg = (*MsgSwapSend)(nil)
var _ sdk.Msg = (*MsgAggregateExchangeRatePrevote)(nil)
var _ sdk.Msg = (*MsgAggregateExchangeRateVote)(nil)
var _ sdk.Msg = (*MsgDelegateFeedConsent)(nil)
var _ sdk.Msg = (*MsgStoreCode)(nil)
var _ sdk.Msg = (*MsgInstantiateContract)(nil)
var _ sdk.Msg = (*MsgExecuteContract)(nil)
var _ sdk.Msg = (*MsgMigrateContract)(nil)
var _ sdk.Msg = (*MsgUpdateContractAdmin)(nil)
var _ sdk.Msg = (*MsgClearContractAdmin)(nil)

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

type MsgSwapSend struct {
	FromAddress sdk.AccAddress `json:"from_address" yaml:"from_address"` // Address of the offer coin payer
	ToAddress   sdk.AccAddress `json:"to_address" yaml:"to_address"`     // Address of the recipient
	OfferCoin   sdk.Coin       `json:"offer_coin" yaml:"offer_coin"`     // Coin being offered
	AskDenom    string         `json:"ask_denom" yaml:"ask_denom"`       // Denom of the coin to swap to
}

type MsgAggregateExchangeRatePrevote struct {
	Hash      string         `json:"hash" yaml:"hash"`
	Feeder    sdk.AccAddress `json:"feeder" yaml:"feeder"`
	Validator sdk.ValAddress `json:"validator" yaml:"validator"`
}

type MsgAggregateExchangeRateVote struct {
	Salt          string         `json:"salt" yaml:"salt"`
	ExchangeRates string         `json:"exchange_rates" yaml:"exchange_rates"` // comma separated dec coins
	Feeder        sdk.AccAddress `json:"feeder" yaml:"feeder"`
	Validator     sdk.ValAddress `json:"validator" yaml:"validator"`
}

type MsgDelegateFeedConsent struct {
	Operator sdk.ValAddress `json:"operator"`
	Delegate sdk.AccAddress `json:"delegate"`
}

type MsgStoreCode struct {
	Sender sdk.AccAddress `json:"sender" yaml:"sender"`
	// WASMByteCode can be raw or gzip compressed
	WASMByteCode string `json:"wasm_byte_code" yaml:"wasm_byte_code"`
}

type MsgInstantiateContract struct {
	Sender    sdk.AccAddress `json:"sender" yaml:"sender"`
	Admin     sdk.AccAddress `json:"admin,omitempty" yaml:"admin,omitempty"`
	CodeID    uint64         `json:"code_id" yaml:"code_id"`
	InitMsg   string         `json:"init_msg" yaml:"init_msg"`
	InitCoins sdk.Coins      `json:"init_coins" yaml:"init_coins"`
}

type MsgExecuteContract struct {
	Sender     sdk.AccAddress `json:"sender" yaml:"sender"`
	Contract   sdk.AccAddress `json:"contract" yaml:"contract"`
	ExecuteMsg string         `json:"execute_msg" yaml:"execute_msg"`
	Coins      sdk.Coins      `json:"coins" yaml:"coins"`
}

type MsgMigrateContract struct {
	Admin      sdk.AccAddress `json:"admin" yaml:"admin"`
	Contract   sdk.AccAddress `json:"contract" yaml:"contract"`
	NewCodeID  uint64         `json:"new_code_id" yaml:"new_code_id"`
	MigrateMsg string         `json:"migrate_msg" yaml:"migrate_msg"`
}

type MsgUpdateContractAdmin struct {
	Admin    sdk.AccAddress `json:"admin" yaml:"admin"`
	NewAdmin sdk.AccAddress `json:"new_admin" yaml:"new_admin"`
	Contract sdk.AccAddress `json:"contract" yaml:"contract"`
}

type MsgClearContractAdmin struct {
	Admin    sdk.AccAddress `json:"admin" yaml:"admin"`
	Contract sdk.AccAddress `json:"contract" yaml:"contract"`
}

type MsgGrantAuthorization struct {
	Granter sdk.AccAddress `json:"granter"`
	Grantee sdk.AccAddress `json:"grantee"`
	Grant   Grant          `json:"grant"`
}

type MsgRevokeAuthorization struct {
	Granter    sdk.AccAddress `json:"granter"`
	Grantee    sdk.AccAddress `json:"grantee"`
	MsgTypeUrl string         `json:"msg_type_url"`
}

type MsgExecAuthorized struct {
	Grantee sdk.AccAddress `json:"grantee"`
	Msgs    []sdk.Msg      `json:"msgs"`
}

type Grant struct {
	Authorization AuthorizationI `json:"authorization"`
	Expiration    time.Time      `json:"expiration"`
}

type AuthorizationI interface{}

var _ AuthorizationI = (*SendAuthorization)(nil)
var _ AuthorizationI = (*GenericAuthorization)(nil)

type SendAuthorization struct {
	SpendLimit sdk.Coins `json:"spend_limit"`
}
type GenericAuthorization struct {
	Msg string `json:"msg"`
}

type MsgGrantAllowance struct {
	// granter is the address of the user granting an allowance of their funds.
	Granter sdk.AccAddress `json:"granter"`
	// grantee is the address of the user being granted an allowance of another user's funds.
	Grantee sdk.AccAddress `json:"grantee"`
	// allowance can be any of basic and filtered fee allowance.
	Allowance FeeAllowanceI `json:"allowance"`
}

type MsgRevokeAllowance struct {
	// granter is the address of the user granting an allowance of their funds.
	Granter sdk.AccAddress `json:"granter"`
	// grantee is the address of the user being granted an allowance of another user's funds.
	Grantee sdk.AccAddress `json:"grantee"`
}

type FeeAllowanceI interface{}

type BasicAllowance struct {
	SpendLimit sdk.Coins `json:"spend_limit"`
	// expiration specifies an optional time when this allowance expires
	Expiration *time.Time `json:"expiration,omitempty"`
}

type PeriodicAllowance struct {
	// basic specifies a struct of `BasicAllowance`
	Basic BasicAllowance `json:"basic"`
	// period specifies the time duration in which period_spend_limit coins can
	// be spent before that allowance is reset
	Period time.Duration `json:"period"`
	// period_spend_limit specifies the maximum number of coins that can be spent
	// in the period
	PeriodSpendLimit sdk.Coins `json:"period_spend_limit"`
	// period_can_spend is the number of coins left to be spent before the period_reset time
	PeriodCanSpend sdk.Coins `json:"period_can_spend"`
	// period_reset is the time at which this period resets and a new one begins,
	// it is calculated from the start time of the first transaction after the
	// last period ended
	PeriodReset time.Time `json:"period_reset"`
}

type AllowedMsgAllowance struct {
	// allowance can be any of basic and filtered fee allowance.
	Allowance FeeAllowanceI `json:"allowance"`
	// allowed_messages are the messages for which the grantee has the access.
	AllowedMessages []string `json:"allowed_messages,omitempty"`
}

const (
	TerraMsgSwap                         = "market/MsgSwap"
	TerraMsgSwapSend                     = "market/MsgSwapSend"
	TerraMsgExchangeRateVote             = "oracle/MsgExchangeRateVote"
	TerraMsgExchangeRatePrevote          = "oracle/MsgExchangeRatePrevote"
	TerraMsgDelegateFeedConsent          = "oracle/MsgDelegateFeedConsent"
	TerraMsgAggregateExchangeRatePrevote = "oracle/MsgAggregateExchangeRatePrevote"
	TerraMsgAggregateExchangeRateVote    = "oracle/MsgAggregateExchangeRateVote"
	TerraSchedule                        = "core/Schedule"
	TerraVestingSchedule                 = "core/VestingSchedule"
	TerraBaseLazyGradedVestingAccount    = "core/LazyGradedVestingAccount"
	TerraMsgStoreCode                    = "wasm/MsgStoreCode"
	TerraMsgInstantiateContract          = "wasm/MsgInstantiateContract"
	TerraMsgExecuteContract              = "wasm/MsgExecuteContract"
	TerraMsgMigrateContract              = "wasm/MsgMigrateContract"
	TerraMsgUpdateContractAdmin          = "wasm/MsgUpdateContractAdmin"
	TerraMsgClearContractAdmin           = "wasm/MsgClearContractAdmin"
	TerraMsgGrantAuthorization           = "msgauth/MsgGrantAuthorization"
	TerraMsgRevokeAuthorization          = "msgauth/MsgRevokeAuthorization"
	TerraMsgExecAuthorized               = "msgauth/MsgExecAuthorized"
	TerraSendAuthorization               = "msgauth/SendAuthorization"
	TerraGenericAuthorization            = "msgauth/GenericAuthorization"
	TerraMsgGrantAllowance               = "feegrant/MsgGrantAllowance"
	TerraMsgRevokeAllowance              = "feegrant/MsgRevokeAllowance"
	TerraBasicAllowance                  = "feegrant/BasicAllowance"
	TerraPeriodicAllowance               = "feegrant/PeriodicAllowance"
	TerraAllowedMsgAllowance             = "feegrant/AllowedMsgAllowance"
)
