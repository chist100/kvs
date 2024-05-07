package keeper

import (
	"kvs/x/kvs/types"
)

var _ types.QueryServer = Keeper{}
