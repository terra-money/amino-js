package types

import (
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
)

type Vote struct {
	Voter      sdk.AccAddress       `json:"voter"`                                                                                          //  address of the voter
	ProposalID uint64               `json:"proposal_id"`                                                                                    //  proposalID of the proposal
	Option     VoteOption           `protobuf:"varint,3,opt,name=option,proto3,enum=cosmos.gov.v1beta1.VoteOption" json:"option,omitempty"` // Deprecated: Do not use.
	Options    []WeightedVoteOption `protobuf:"bytes,4,rep,name=options,proto3" json:"options"`
}

type Votes []Vote

type VoteOption int32
