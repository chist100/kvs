package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"kvs/x/kvs/types"
)

func (k msgServer) DataConfirmation(goCtx context.Context, msg *types.MsgDataConfirmation) (*types.MsgDataConfirmationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDataConfirmationResponse{}, nil
}
