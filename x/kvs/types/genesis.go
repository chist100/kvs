package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DataList:     []Data{},
		ProposalList: []Proposal{},
		Acl:          nil,
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in data
	dataIndexMap := make(map[string]struct{})

	for _, elem := range gs.DataList {
		index := string(DataKey(elem.Index))
		if _, ok := dataIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for data")
		}
		dataIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in proposal
	proposalIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProposalList {
		index := string(ProposalKey(elem.Index))
		if _, ok := proposalIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for proposal")
		}
		proposalIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
