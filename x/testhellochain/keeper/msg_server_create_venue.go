package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"testhellochain/x/testhellochain/types"
)

func (k msgServer) CreateVenue(goCtx context.Context, msg *types.MsgCreateVenue) (*types.MsgCreateVenueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	// Create variable of type Post
	var venue = types.Venue{
		Creator:  msg.Creator,
		Category: msg.Category,
		Name:     msg.Name,
	}

	// Add a post to the store and get back the ID
	id := k.AppendVenue(ctx, venue)

	// Return the ID of the post

	return &types.MsgCreateVenueResponse{Id: id}, nil
}
