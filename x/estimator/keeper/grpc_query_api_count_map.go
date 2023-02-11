package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testhellochain/x/estimator/types"
)

func (k Keeper) ApiCountMapAll(c context.Context, req *types.QueryAllApiCountMapRequest) (*types.QueryAllApiCountMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var apiCountMaps []types.ApiCountMap
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	apiCountMapStore := prefix.NewStore(store, types.KeyPrefix(types.ApiCountMapKeyPrefix))

	pageRes, err := query.Paginate(apiCountMapStore, req.Pagination, func(key []byte, value []byte) error {
		var apiCountMap types.ApiCountMap
		if err := k.cdc.Unmarshal(value, &apiCountMap); err != nil {
			return err
		}

		apiCountMaps = append(apiCountMaps, apiCountMap)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllApiCountMapResponse{ApiCountMap: apiCountMaps, Pagination: pageRes}, nil
}

func (k Keeper) ApiCountMap(c context.Context, req *types.QueryGetApiCountMapRequest) (*types.QueryGetApiCountMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetApiCountMap(
		ctx,
		req.Creator,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetApiCountMapResponse{ApiCountMap: val}, nil
}
