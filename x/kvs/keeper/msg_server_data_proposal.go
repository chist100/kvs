package keeper

import (
	"context"
	"fmt"

	"kvs/x/kvs/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) DataProposal(goCtx context.Context, msg *types.MsgDataProposal) (*types.MsgDataProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, exist := k.Keeper.GetProposal(ctx, msg.Key); exist {
		k.Logger(ctx).Error("proposal already exist, input %s", msg)
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("proposal with this key: %s already exist", msg.Key))
	}

	if _, exist := k.Keeper.GetData(ctx, msg.Key); exist {
		k.Logger(ctx).Error("data already exist, input %s", msg)
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("data with this key: %s already exist", msg.Key))
	}

	k.Keeper.SetProposal(ctx, types.Proposal{
		Index: msg.Key,
		Value: msg.Value,
		Acknowledgments: []string{},
	})
	k.Logger(ctx).Info("proposal successfully stored, input %s", msg)

	return &types.MsgDataProposalResponse{}, nil
}
