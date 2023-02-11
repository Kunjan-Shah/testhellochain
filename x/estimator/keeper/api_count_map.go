package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testhellochain/x/estimator/types"
)

// SetApiCountMap set a specific apiCountMap in the store from its index
func (k Keeper) SetApiCountMap(ctx sdk.Context, apiCountMap types.ApiCountMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApiCountMapKeyPrefix))
	b := k.cdc.MustMarshal(&apiCountMap)
	store.Set(types.ApiCountMapKey(
		apiCountMap.Creator,
	), b)
}

// GetApiCountMap returns a apiCountMap from its index
func (k Keeper) GetApiCountMap(
	ctx sdk.Context,
	creator string,

) (val types.ApiCountMap, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApiCountMapKeyPrefix))

	b := store.Get(types.ApiCountMapKey(
		creator,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveApiCountMap removes a apiCountMap from the store
func (k Keeper) RemoveApiCountMap(
	ctx sdk.Context,
	creator string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApiCountMapKeyPrefix))
	store.Delete(types.ApiCountMapKey(
		creator,
	))
}

// GetAllApiCountMap returns all apiCountMap
func (k Keeper) GetAllApiCountMap(ctx sdk.Context) (list []types.ApiCountMap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApiCountMapKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ApiCountMap
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
