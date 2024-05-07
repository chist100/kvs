package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDataConfirmation = "data_confirmation"

var _ sdk.Msg = &MsgDataConfirmation{}

func NewMsgDataConfirmation(creator string, key string) *MsgDataConfirmation {
	return &MsgDataConfirmation{
		Creator: creator,
		Key:     key,
	}
}

func (msg *MsgDataConfirmation) Route() string {
	return RouterKey
}

func (msg *MsgDataConfirmation) Type() string {
	return TypeMsgDataConfirmation
}

func (msg *MsgDataConfirmation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDataConfirmation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDataConfirmation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
