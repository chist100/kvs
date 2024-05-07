package kvs_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "kvs/testutil/keeper"
	"kvs/testutil/nullify"
	"kvs/x/kvs"
	"kvs/x/kvs/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DataList: []types.Data{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KvsKeeper(t)
	kvs.InitGenesis(ctx, *k, genesisState)
	got := kvs.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DataList, got.DataList)
	// this line is used by starport scaffolding # genesis/test/assert
}
