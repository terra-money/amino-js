package extensions

import (
	"time"

	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
	"github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/x/auth"
)

var _ sdk.Msg = (*MsgSwap)(nil)
var _ sdk.Msg = (*MsgSwapSend)(nil)
var _ sdk.Msg = (*MsgExchangeRatePrevote)(nil)
var _ sdk.Msg = (*MsgExchangeRateVote)(nil)
var _ sdk.Msg = (*MsgAggregateExchangeRatePrevote)(nil)
var _ sdk.Msg = (*MsgAggregateExchangeRateVote)(nil)
var _ sdk.Msg = (*MsgDelegateFeedConsent)(nil)
var _ sdk.Msg = (*MsgStoreCode)(nil)
var _ sdk.Msg = (*MsgInstantiateContract)(nil)
var _ sdk.Msg = (*MsgExecuteContract)(nil)
var _ sdk.Msg = (*MsgMigrateContract)(nil)
var _ sdk.Msg = (*MsgUpdateContractOwner)(nil)

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
	Owner      sdk.AccAddress `json:"owner" yaml:"owner"`
	CodeID     uint64         `json:"code_id" yaml:"code_id"`
	InitMsg    string         `json:"init_msg" yaml:"init_msg"`
	InitCoins  sdk.Coins      `json:"init_coins" yaml:"init_coins"`
	Migratable bool           `json:"migratable" yaml:"migratable"`
}

type MsgExecuteContract struct {
	Sender     sdk.AccAddress `json:"sender" yaml:"sender"`
	Contract   sdk.AccAddress `json:"contract" yaml:"contract"`
	ExecuteMsg string         `json:"execute_msg" yaml:"execute_msg"`
	Coins      sdk.Coins      `json:"coins" yaml:"coins"`
}

type MsgMigrateContract struct {
	Owner      sdk.AccAddress `json:"owner" yaml:"owner"`
	Contract   sdk.AccAddress `json:"contract" yaml:"contract"`
	NewCodeID  uint64         `json:"new_code_id" yaml:"new_code_id"`
	MigrateMsg string         `json:"migrate_msg" yaml:"migrate_msg"`
}

type MsgUpdateContractOwner struct {
	Owner    sdk.AccAddress `json:"owner" yaml:"owner"`
	NewOwner sdk.AccAddress `json:"new_owner" yaml:"new_owner"`
	Contract sdk.AccAddress `json:"contract" yaml:"contract"`
}

type MsgGrantAuthorization struct {
	Granter       sdk.AccAddress `json:"granter"`
	Grantee       sdk.AccAddress `json:"grantee"`
	Authorization Authorization  `json:"authorization"`
	Period        time.Duration  `json:"period"`
}

type MsgRevokeAuthorization struct {
	Granter              sdk.AccAddress `json:"granter"`
	Grantee              sdk.AccAddress `json:"grantee"`
	AuthorizationMsgType string         `json:"authorization_msg_type"`
}

type MsgExecAuthorized struct {
	Grantee sdk.AccAddress `json:"grantee"`
	Msgs    []sdk.Msg      `json:"msgs"`
}

type Authorization interface{}

var _ Authorization = (*SendAuthorization)(nil)
var _ Authorization = (*GenericAuthorization)(nil)

type SendAuthorization struct {
	SpendLimit sdk.Coins `json:"spend_limit"`
}
type GenericAuthorization struct {
	GrantMsgType string `json:"grant_msg_type"`
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
	TerraMsgUpdateContractOwner          = "wasm/MsgUpdateContractOwner"
	TerraMsgGrantAuthorization           = "msgauth/MsgGrantAuthorization"
	TerraMsgRevokeAuthorization          = "msgauth/MsgRevokeAuthorization"
	TerraMsgExecAuthorized               = "msgauth/MsgExecAuthorized"
	TerraSendAuthorization               = "msgauth/SendAuthorization"
	TerraGenericAuthorization            = "msgauth/GenericAuthorization"
)
