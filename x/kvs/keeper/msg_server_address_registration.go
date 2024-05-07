package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"kvs/x/kvs/types"
)

func (k msgServer) AddressRegistration(goCtx context.Context, msg *types.MsgAddressRegistration) (*types.MsgAddressRegistrationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddressRegistrationResponse{}, nil
}
