package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "kvs/testutil/keeper"
	"kvs/testutil/nullify"
	"kvs/x/kvs/keeper"
	"kvs/x/kvs/types"
)

func createTestAcl(keeper *keeper.Keeper, ctx sdk.Context) types.Acl {
	item := types.Acl{}
	keeper.SetAcl(ctx, item)
	return item
}

func TestAclGet(t *testing.T) {
	keeper, ctx := keepertest.KvsKeeper(t)
	item := createTestAcl(keeper, ctx)
	rst, found := keeper.GetAcl(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestAclRemove(t *testing.T) {
	keeper, ctx := keepertest.KvsKeeper(t)
	createTestAcl(keeper, ctx)
	keeper.RemoveAcl(ctx)
	_, found := keeper.GetAcl(ctx)
	require.False(t, found)
}
