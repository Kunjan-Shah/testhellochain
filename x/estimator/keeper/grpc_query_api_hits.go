package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testhellochain/x/estimator/types"
)

func (k Keeper) ApiHitsAll(c context.Context, req *types.QueryAllApiHitsRequest) (*types.QueryAllApiHitsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var apiHitss []types.ApiHits
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	apiHitsStore := prefix.NewStore(store, types.KeyPrefix(types.ApiHitsKey))

	pageRes, err := query.Paginate(apiHitsStore, req.Pagination, func(key []byte, value []byte) error {
		var apiHits types.ApiHits
		if err := k.cdc.Unmarshal(value, &apiHits); err != nil {
			return err
		}

		apiHitss = append(apiHitss, apiHits)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllApiHitsResponse{ApiHits: apiHitss, Pagination: pageRes}, nil
}

func (k Keeper) ApiHits(c context.Context, req *types.QueryGetApiHitsRequest) (*types.QueryGetApiHitsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	apiHits, found := k.GetApiHits(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetApiHitsResponse{ApiHits: apiHits}, nil
}
