package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
)

type Proposal struct {
	Content `json:"content"` // Proposal content interface

	ProposalID       uint64         `json:"id"`                 //  ID of the proposal
	Status           ProposalStatus `json:"proposal_status"`    // Status of the Proposal {Pending, Active, Passed, Rejected}
	FinalTallyResult TallyResult    `json:"final_tally_result"` // Result of Tallys

	SubmitTime     time.Time `json:"submit_time"`      // Time of the block where TxGovSubmitProposal was included
	DepositEndTime time.Time `json:"deposit_end_time"` // Time that the Proposal would expire if deposit amount isn't met
	TotalDeposit   sdk.Coins `json:"total_deposit"`    // Current deposit on this proposal. Initial value is set at InitialDeposit

	VotingStartTime time.Time `json:"voting_start_time"` // Time of the block where MinDeposit was reached. -1 if MinDeposit is not reached
	VotingEndTime   time.Time `json:"voting_end_time"`   // Time that the VotingPeriod for this proposal will end and votes will be tallied
}

func (p Proposal) String() string {
	return fmt.Sprintf(`Proposal %d:
  Title:              %s
  Type:               %s
  Status:             %s
  Submit Time:        %s
  Deposit End Time:   %s
  Total Deposit:      %s
  Voting Start Time:  %s
  Voting End Time:    %s
  Description:        %s`,
		p.ProposalID, p.GetTitle(), p.ProposalType(),
		p.Status, p.SubmitTime, p.DepositEndTime,
		p.TotalDeposit, p.VotingStartTime, p.VotingEndTime, p.GetDescription(),
	)
}

type Proposals []Proposal

func (p Proposals) String() string {
	out := "ID - (Status) [Type] Title\n"
	for _, prop := range p {
		out += fmt.Sprintf("%d - (%s) [%s] %s\n",
			prop.ProposalID, prop.Status,
			prop.ProposalType(), prop.GetTitle())
	}
	return strings.TrimSpace(out)
}

type (
	ProposalQueue  []uint64
	ProposalStatus int32
)

var ProposalStatus_name = map[int32]string{
	0: "PROPOSAL_STATUS_UNSPECIFIED",
	1: "PROPOSAL_STATUS_DEPOSIT_PERIOD",
	2: "PROPOSAL_STATUS_VOTING_PERIOD",
	3: "PROPOSAL_STATUS_PASSED",
	4: "PROPOSAL_STATUS_REJECTED",
	5: "PROPOSAL_STATUS_FAILED",
}

var ProposalStatus_value = map[string]int32{
	"PROPOSAL_STATUS_UNSPECIFIED":    0,
	"PROPOSAL_STATUS_DEPOSIT_PERIOD": 1,
	"PROPOSAL_STATUS_VOTING_PERIOD":  2,
	"PROPOSAL_STATUS_PASSED":         3,
	"PROPOSAL_STATUS_REJECTED":       4,
	"PROPOSAL_STATUS_FAILED":         5,
}

func (status ProposalStatus) Marshal() ([]byte, error) {
	return []byte{byte(status)}, nil
}

func (status *ProposalStatus) Unmarshal(data []byte) error {
	*status = ProposalStatus(data[0])
	return nil
}

func (status ProposalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(status.String())
}

func (status *ProposalStatus) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	stat, ok := ProposalStatus_value[s]
	if !ok {
		return fmt.Errorf("'%s' is not a valid proposal status", s)
	}

	*status = ProposalStatus(stat)
	return nil
}

func (status ProposalStatus) String() string {
	return ProposalStatus_name[int32(status)]
}

type TallyResult struct {
	Yes        sdk.Int `json:"yes"`
	Abstain    sdk.Int `json:"abstain"`
	No         sdk.Int `json:"no"`
	NoWithVeto sdk.Int `json:"no_with_veto"`
}

func (tr TallyResult) String() string {
	return fmt.Sprintf(`Tally Result:
  Yes:        %s
  Abstain:    %s
  No:         %s
  NoWithVeto: %s`, tr.Yes, tr.Abstain, tr.No, tr.NoWithVeto)
}

const (
	ProposalTypeText            string = "Text"
	ProposalTypeSoftwareUpgrade string = "SoftwareUpgrade"
)

type TextProposal struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var _ Content = TextProposal{}

func (tp TextProposal) GetTitle() string       { return tp.Title }
func (tp TextProposal) GetDescription() string { return tp.Description }
func (tp TextProposal) ProposalType() string   { return ProposalTypeText }

func (tp TextProposal) String() string {
	return fmt.Sprintf(`Text Proposal:
  Title:       %s
  Description: %s
`, tp.Title, tp.Description)
}

type SoftwareUpgradeProposal struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var _ Content = SoftwareUpgradeProposal{}

func (sup SoftwareUpgradeProposal) GetTitle() string       { return sup.Title }
func (sup SoftwareUpgradeProposal) GetDescription() string { return sup.Description }
func (sup SoftwareUpgradeProposal) ProposalType() string   { return ProposalTypeSoftwareUpgrade }

func (sup SoftwareUpgradeProposal) String() string {
	return fmt.Sprintf(`Software Upgrade Proposal:
  Title:       %s
  Description: %s
`, sup.Title, sup.Description)
}
