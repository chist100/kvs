package keeper

import (
	"context"
	"fmt"
	"strings"

	"kvs/x/kvs/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) AddressRegistration(goCtx context.Context, msg *types.MsgAddressRegistration) (*types.MsgAddressRegistrationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, exist := k.Keeper.GetAcl(ctx); exist {
		k.Logger(ctx).Error("addresses already registered, input %s", msg)
		return nil, status.Error(codes.AlreadyExists, "addresses already registered")
	}
	if len(msg.Addresses) != 3 {
		k.Logger(ctx).Error("addresses array length is not equal three, input %s", msg)
		return nil, status.Error(codes.InvalidArgument, "addresses array length is not equal three")
	}

	for _, addr := range msg.Addresses {
		if len(strings.TrimSpace(addr)) == 0 {
			k.Logger(ctx).Error("empty address string is not allowed, input %s", msg)
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("empty address string is not allowed"))
		}
	}

	k.Keeper.SetAcl(ctx, types.Acl{
		Adresses: msg.Addresses,
	})
	k.Logger(ctx).Info("address successfully registered, input %s", msg)
	return &types.MsgAddressRegistrationResponse{}, nil
}
