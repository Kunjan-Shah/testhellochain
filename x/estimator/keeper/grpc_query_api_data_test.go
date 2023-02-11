package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "testhellochain/testutil/keeper"
	"testhellochain/testutil/nullify"
	"testhellochain/x/estimator/types"
)

func TestApiDataQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EstimatorKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNApiData(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetApiDataRequest
		response *types.QueryGetApiDataResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetApiDataRequest{Id: msgs[0].Id},
			response: &types.QueryGetApiDataResponse{ApiData: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetApiDataRequest{Id: msgs[1].Id},
			response: &types.QueryGetApiDataResponse{ApiData: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetApiDataRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ApiData(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestApiDataQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EstimatorKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNApiData(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllApiDataRequest {
		return &types.QueryAllApiDataRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ApiDataAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ApiData), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ApiData),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ApiDataAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ApiData), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ApiData),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ApiDataAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ApiData),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ApiDataAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
