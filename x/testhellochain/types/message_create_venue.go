package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateVenue = "create_venue"

var _ sdk.Msg = &MsgCreateVenue{}

func NewMsgCreateVenue(creator string, category string, name string) *MsgCreateVenue {
	return &MsgCreateVenue{
		Creator:  creator,
		Category: category,
		Name:     name,
	}
}

func (msg *MsgCreateVenue) Route() string {
	return RouterKey
}

func (msg *MsgCreateVenue) Type() string {
	return TypeMsgCreateVenue
}

func (msg *MsgCreateVenue) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateVenue) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVenue) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
