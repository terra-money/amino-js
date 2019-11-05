package extensions

import (
	amino "github.com/tendermint/go-amino"
)

func RegisterCodec(codec *amino.Codec) {
	codec.RegisterConcrete(MsgExchangeRateVote{}, TerraMsgExchangeRateVote, nil)
	codec.RegisterConcrete(MsgExchangeRatePrevote{}, TerraMsgExchangeRatePrevote, nil)
	codec.RegisterConcrete(MsgDelegateFeedConsent{}, TerraMsgDelegateFeedConsent, nil)

	codec.RegisterConcrete(MsgSwap{}, TerraMsgSwap, nil)

	codec.RegisterConcrete(&Schedule{}, TerraSchedule, nil)
	codec.RegisterConcrete(&VestingSchedule{}, TerraVestingSchedule, nil)
	codec.RegisterConcrete(&BaseLazyGradedVestingAccount{}, TerraBaseLazyGradedVestingAccount, nil)
}
