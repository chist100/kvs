package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddressRegistration = "address_registration"

var _ sdk.Msg = &MsgAddressRegistration{}

func NewMsgAddressRegistration(creator string, addresses []string) *MsgAddressRegistration {
	return &MsgAddressRegistration{
		Creator:   creator,
		Addresses: addresses,
	}
}

func (msg *MsgAddressRegistration) Route() string {
	return RouterKey
}

func (msg *MsgAddressRegistration) Type() string {
	return TypeMsgAddressRegistration
}

func (msg *MsgAddressRegistration) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddressRegistration) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddressRegistration) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
