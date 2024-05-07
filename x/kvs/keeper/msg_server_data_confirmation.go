package keeper

import (
	"context"
	"fmt"

	"kvs/x/kvs/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) DataConfirmation(goCtx context.Context, msg *types.MsgDataConfirmation) (*types.MsgDataConfirmationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	acl, exist := k.Keeper.GetAcl(ctx)
	if !exist {
		k.Logger(ctx).Error("address is not registered yet, input %s", msg)
		return nil, status.Error(codes.Internal, fmt.Sprintf("address is not registered yet"))
	}

	if !k.Keeper.CheckAddress(ctx, msg.Creator) {
		k.Logger(ctx).Error("address is not existence in acl store, input %s", msg)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("address '%s' is not existence in acl store", msg.Creator))
	}

	prop, exist := k.Keeper.GetProposal(ctx, msg.Key)
	if !exist {
		k.Logger(ctx).Error("proposal is not exist, input %s", msg)
		return nil, status.Error(codes.NotFound, "proposal with key: %s is not exist")
	}

	for _, ack := range prop.Acknowledgments {
		if msg.Creator == ack {
			k.Logger(ctx).Error("address is already confirm proposal, input %s", msg)
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("address '%s' is already confirm proposal  with key: %s", msg.Creator, msg.Key))
		}
	}

	prop.Acknowledgments = append(prop.Acknowledgments, msg.Creator)
	k.Keeper.SetProposal(ctx, prop)
	k.Logger(ctx).Info("data successfully confirmed, input %s", msg)

	if len(prop.Acknowledgments) == len(acl.Adresses) {
		k.Keeper.SetData(ctx, types.Data{
			Index: prop.Index,
			Value: prop.Value,
		})
		k.Logger(ctx).Info("data successfully stored, after all confirmation, input %s", msg)
	}
	return &types.MsgDataConfirmationResponse{}, nil
}
