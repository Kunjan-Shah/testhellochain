package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEstimate = "estimate"

var _ sdk.Msg = &MsgEstimate{}

func NewMsgEstimate(creator string, start uint64, end uint64) *MsgEstimate {
	return &MsgEstimate{
		Creator: creator,
		Start:   start,
		End:     end,
	}
}

func (msg *MsgEstimate) Route() string {
	return RouterKey
}

func (msg *MsgEstimate) Type() string {
	return TypeMsgEstimate
}

func (msg *MsgEstimate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEstimate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEstimate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
