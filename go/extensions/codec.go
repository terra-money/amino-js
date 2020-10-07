package extensions

import (
	amino "github.com/tendermint/go-amino"
)

func RegisterCodec(codec *amino.Codec) {
	codec.RegisterConcrete(MsgExchangeRateVote{}, TerraMsgExchangeRateVote, nil)
	codec.RegisterConcrete(MsgExchangeRatePrevote{}, TerraMsgExchangeRatePrevote, nil)
	codec.RegisterConcrete(MsgAggregateExchangeRateVote{}, TerraMsgAggregateExchangeRateVote, nil)
	codec.RegisterConcrete(MsgAggregateExchangeRatePrevote{}, TerraMsgAggregateExchangeRatePrevote, nil)
	codec.RegisterConcrete(MsgDelegateFeedConsent{}, TerraMsgDelegateFeedConsent, nil)

	codec.RegisterConcrete(MsgSwap{}, TerraMsgSwap, nil)
	codec.RegisterConcrete(MsgSwapSend{}, TerraMsgSwapSend, nil)

	codec.RegisterConcrete(MsgStoreCode{}, TerraMsgStoreCode, nil)
	codec.RegisterConcrete(MsgInstantiateContract{}, TerraMsgInstantiateContract, nil)
	codec.RegisterConcrete(MsgExecuteContract{}, TerraMsgExecuteContract, nil)
	codec.RegisterConcrete(MsgMigrateContract{}, TerraMsgMigrateContract, nil)
	codec.RegisterConcrete(MsgUpdateContractOwner{}, TerraMsgUpdateContractOwner, nil)

	codec.RegisterConcrete(&Schedule{}, TerraSchedule, nil)
	codec.RegisterConcrete(&VestingSchedule{}, TerraVestingSchedule, nil)
	codec.RegisterConcrete(&BaseLazyGradedVestingAccount{}, TerraBaseLazyGradedVestingAccount, nil)

	codec.RegisterConcrete(MsgGrantAuthorization{}, TerraMsgGrantAuthorization, nil)
	codec.RegisterConcrete(MsgRevokeAuthorization{}, TerraMsgRevokeAuthorization, nil)
	codec.RegisterConcrete(MsgExecAuthorized{}, TerraMsgExecAuthorized, nil)
	codec.RegisterConcrete(SendAuthorization{}, TerraSendAuthorization, nil)
	codec.RegisterConcrete(GenericAuthorization{}, TerraGenericAuthorization, nil)

	codec.RegisterInterface((*Authorization)(nil), nil)
}
