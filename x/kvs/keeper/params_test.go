package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "kvs/testutil/keeper"
	"kvs/x/kvs/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.KvsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
