package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"kvs/x/kvs/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				DataList: []types.Data{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				ProposalList: []types.Proposal{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				Acl: &types.Acl{
					Adresses: []string{"15"},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated data",
			genState: &types.GenesisState{
				DataList: []types.Data{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated proposal",
			genState: &types.GenesisState{
				ProposalList: []types.Proposal{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
