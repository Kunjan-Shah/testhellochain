package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testhellochain/x/estimator/types"
)

// GetApiDataCount get the total number of apiData
func (k Keeper) GetApiDataCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ApiDataCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetApiDataCount set the total number of apiData
func (k Keeper) SetApiDataCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ApiDataCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendApiData appends a apiData in the store with a new id and update the count
func (k Keeper) AppendApiData(
	ctx sdk.Context,
	apiData types.ApiData,
) uint64 {
	// Create the apiData
	count := k.GetApiDataCount(ctx)

	// Set the ID of the appended value
	apiData.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApiDataKey))
	appendedValue := k.cdc.MustMarshal(&apiData)
	store.Set(GetApiDataIDBytes(apiData.Id), appendedValue)

	// Update apiData count
	k.SetApiDataCount(ctx, count+1)

	return count
}

// SetApiData set a specific apiData in the store
func (k Keeper) SetApiData(ctx sdk.Context, apiData types.ApiData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApiDataKey))
	b := k.cdc.MustMarshal(&apiData)
	store.Set(GetApiDataIDBytes(apiData.Id), b)
}

// GetApiData returns a apiData from its id
func (k Keeper) GetApiData(ctx sdk.Context, id uint64) (val types.ApiData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApiDataKey))
	b := store.Get(GetApiDataIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveApiData removes a apiData from the store
func (k Keeper) RemoveApiData(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApiDataKey))
	store.Delete(GetApiDataIDBytes(id))
}

// GetAllApiData returns all apiData
func (k Keeper) GetAllApiData(ctx sdk.Context) (list []types.ApiData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApiDataKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ApiData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetApiDataIDBytes returns the byte representation of the ID
func GetApiDataIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetApiDataIDFromBytes returns ID in uint64 format from a byte array
func GetApiDataIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
