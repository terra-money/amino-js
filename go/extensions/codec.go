package extensions

import (
	amino "github.com/tendermint/go-amino"
)

func RegisterCodec(codec *amino.Codec) {
	codec.RegisterConcrete(MsgAggregateExchangeRateVote{}, TerraMsgAggregateExchangeRateVote, nil)
	codec.RegisterConcrete(MsgAggregateExchangeRatePrevote{}, TerraMsgAggregateExchangeRatePrevote, nil)
	codec.RegisterConcrete(MsgDelegateFeedConsent{}, TerraMsgDelegateFeedConsent, nil)

	codec.RegisterConcrete(MsgSwap{}, TerraMsgSwap, nil)
	codec.RegisterConcrete(MsgSwapSend{}, TerraMsgSwapSend, nil)

	codec.RegisterConcrete(MsgStoreCode{}, TerraMsgStoreCode, nil)
	codec.RegisterConcrete(MsgInstantiateContract{}, TerraMsgInstantiateContract, nil)
	codec.RegisterConcrete(MsgExecuteContract{}, TerraMsgExecuteContract, nil)
	codec.RegisterConcrete(MsgMigrateContract{}, TerraMsgMigrateContract, nil)
	codec.RegisterConcrete(MsgUpdateContractAdmin{}, TerraMsgUpdateContractAdmin, nil)
	codec.RegisterConcrete(MsgClearContractAdmin{}, TerraMsgClearContractAdmin, nil)

	codec.RegisterConcrete(&Schedule{}, TerraSchedule, nil)
	codec.RegisterConcrete(&VestingSchedule{}, TerraVestingSchedule, nil)
	codec.RegisterConcrete(&BaseLazyGradedVestingAccount{}, TerraBaseLazyGradedVestingAccount, nil)

	codec.RegisterConcrete(MsgGrantAuthorization{}, TerraMsgGrantAuthorization, nil)
	codec.RegisterConcrete(MsgRevokeAuthorization{}, TerraMsgRevokeAuthorization, nil)
	codec.RegisterConcrete(MsgExecAuthorized{}, TerraMsgExecAuthorized, nil)
	codec.RegisterConcrete(SendAuthorization{}, TerraSendAuthorization, nil)
	codec.RegisterConcrete(GenericAuthorization{}, TerraGenericAuthorization, nil)

	codec.RegisterConcrete(BasicAllowance{}, TerraBasicAllowance, nil)
	codec.RegisterConcrete(PeriodicAllowance{}, TerraPeriodicAllowance, nil)
	codec.RegisterConcrete(AllowedMsgAllowance{}, TerraAllowedMsgAllowance, nil)
	codec.RegisterConcrete(MsgGrantAllowance{}, TerraMsgGrantAllowance, nil)
	codec.RegisterConcrete(MsgRevokeAllowance{}, TerraMsgRevokeAllowance, nil)

	codec.RegisterInterface((*AuthorizationI)(nil), nil)
	codec.RegisterInterface((*FeeAllowanceI)(nil), nil)
}
