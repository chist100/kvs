package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "kvs/testutil/keeper"
	"kvs/testutil/nullify"
	"kvs/x/kvs/keeper"
	"kvs/x/kvs/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNProposal(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Proposal {
	items := make([]types.Proposal, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetProposal(ctx, items[i])
	}
	return items
}

func TestProposalGet(t *testing.T) {
	keeper, ctx := keepertest.KvsKeeper(t)
	items := createNProposal(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProposal(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestProposalRemove(t *testing.T) {
	keeper, ctx := keepertest.KvsKeeper(t)
	items := createNProposal(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProposal(ctx,
			item.Index,
		)
		_, found := keeper.GetProposal(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestProposalGetAll(t *testing.T) {
	keeper, ctx := keepertest.KvsKeeper(t)
	items := createNProposal(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProposal(ctx)),
	)
}
