package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testhellochain/x/estimator/types"
)

// GetApiHitsCount get the total number of apiHits
func (k Keeper) GetApiHitsCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ApiHitsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetApiHitsCount set the total number of apiHits
func (k Keeper) SetApiHitsCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ApiHitsCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendApiHits appends a apiHits in the store with a new id and update the count
func (k Keeper) AppendApiHits(
	ctx sdk.Context,
	apiHits types.ApiHits,
) uint64 {
	// Create the apiHits
	count := k.GetApiHitsCount(ctx)

	// Set the ID of the appended value
	apiHits.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApiHitsKey))
	appendedValue := k.cdc.MustMarshal(&apiHits)
	store.Set(GetApiHitsIDBytes(apiHits.Id), appendedValue)

	// Update apiHits count
	k.SetApiHitsCount(ctx, count+1)

	return count
}

// SetApiHits set a specific apiHits in the store
func (k Keeper) SetApiHits(ctx sdk.Context, apiHits types.ApiHits) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApiHitsKey))
	b := k.cdc.MustMarshal(&apiHits)
	store.Set(GetApiHitsIDBytes(apiHits.Id), b)
}

// GetApiHits returns a apiHits from its id
func (k Keeper) GetApiHits(ctx sdk.Context, id uint64) (val types.ApiHits, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApiHitsKey))
	b := store.Get(GetApiHitsIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveApiHits removes a apiHits from the store
func (k Keeper) RemoveApiHits(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApiHitsKey))
	store.Delete(GetApiHitsIDBytes(id))
}

// GetAllApiHits returns all apiHits
func (k Keeper) GetAllApiHits(ctx sdk.Context) (list []types.ApiHits) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApiHitsKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ApiHits
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetApiHitsIDBytes returns the byte representation of the ID
func GetApiHitsIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetApiHitsIDFromBytes returns ID in uint64 format from a byte array
func GetApiHitsIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
