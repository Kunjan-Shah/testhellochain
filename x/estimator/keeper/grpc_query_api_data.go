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

func (k Keeper) ApiDataAll(c context.Context, req *types.QueryAllApiDataRequest) (*types.QueryAllApiDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var apiDatas []types.ApiData
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	apiDataStore := prefix.NewStore(store, types.KeyPrefix(types.ApiDataKey))

	pageRes, err := query.Paginate(apiDataStore, req.Pagination, func(key []byte, value []byte) error {
		var apiData types.ApiData
		if err := k.cdc.Unmarshal(value, &apiData); err != nil {
			return err
		}

		apiDatas = append(apiDatas, apiData)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllApiDataResponse{ApiData: apiDatas, Pagination: pageRes}, nil
}

func (k Keeper) ApiData(c context.Context, req *types.QueryGetApiDataRequest) (*types.QueryGetApiDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	apiData, found := k.GetApiData(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetApiDataResponse{ApiData: apiData}, nil
}
