package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"kvs/x/kvs/types"
)

func (k msgServer) DataProposal(goCtx context.Context, msg *types.MsgDataProposal) (*types.MsgDataProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDataProposalResponse{}, nil
}
