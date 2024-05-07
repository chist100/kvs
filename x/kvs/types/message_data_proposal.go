package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDataProposal = "data_proposal"

var _ sdk.Msg = &MsgDataProposal{}

func NewMsgDataProposal(creator string, key string, value string) *MsgDataProposal {
	return &MsgDataProposal{
		Creator: creator,
		Key:     key,
		Value:   value,
	}
}

func (msg *MsgDataProposal) Route() string {
	return RouterKey
}

func (msg *MsgDataProposal) Type() string {
	return TypeMsgDataProposal
}

func (msg *MsgDataProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDataProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDataProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
