package extensions

import (
	amino "github.com/tendermint/go-amino"
)

func RegisterCodec(codec *amino.Codec) {
	codec.RegisterConcrete(MsgPriceVote{}, TerraMsgPriceVote, nil)
	codec.RegisterConcrete(MsgPricePrevote{}, TerraMsgPricePrevote, nil)
	codec.RegisterConcrete(MsgDelegateFeederPermission{}, TerraMsgDelegateFeederPermission, nil)

	codec.RegisterConcrete(MsgSwap{}, TerraMsgSwap, nil)

	codec.RegisterConcrete(&Schedule{}, TerraSchedule, nil)
	codec.RegisterConcrete(&VestingSchedule{}, TerraVestingSchedule, nil)
	codec.RegisterConcrete(&BaseLazyGradedVestingAccount{}, TerraBaseLazyGradedVestingAccount, nil)
}
