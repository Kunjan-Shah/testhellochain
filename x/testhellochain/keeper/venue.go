package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"testhellochain/x/testhellochain/types"
)

// func (k Keeper) AppendVenue() uint64 {
//   count := k.GetVenueCount()
//   store.Set()
//   k.SetVenueCount()
//   return count
// }

func (k Keeper) GetVenueCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post/count/")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.VenueCountKey))

	// Convert the PostCountKey to bytes
	byteKey := []byte(types.VenueCountKey)

	// Get the value of the count
	bz := store.Get(byteKey)

	// Return zero if the count value is not found (for example, it's the first post)
	if bz == nil {
		return 0
	}

	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetVenueCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post/count/")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.VenueCountKey))

	// Convert the PostCountKey to bytes
	byteKey := []byte(types.VenueCountKey)

	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	// Set the value of Post/count/ to count
	store.Set(byteKey, bz)
}

func (k Keeper) AppendVenue(ctx sdk.Context, venue types.Venue) uint64 {
	// Get the current number of posts in the store
	count := k.GetVenueCount(ctx)

	// Assign an ID to the post based on the number of posts in the store
	venue.Id = count

	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.VenueKey))

	// Convert the post ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, venue.Id)

	// Marshal the post into bytes
	appendedValue := k.cdc.MustMarshal(&venue)

	// Insert the post bytes using post ID as a key
	store.Set(byteKey, appendedValue)

	// Update the post count
	k.SetVenueCount(ctx, count+1)
	return count
}
