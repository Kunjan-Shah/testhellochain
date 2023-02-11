package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "testhellochain/testutil/keeper"
	"testhellochain/testutil/nullify"
	"testhellochain/x/estimator/keeper"
	"testhellochain/x/estimator/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNApiCountMap(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ApiCountMap {
	items := make([]types.ApiCountMap, n)
	for i := range items {
		items[i].Creator = strconv.Itoa(i)

		keeper.SetApiCountMap(ctx, items[i])
	}
	return items
}

func TestApiCountMapGet(t *testing.T) {
	keeper, ctx := keepertest.EstimatorKeeper(t)
	items := createNApiCountMap(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetApiCountMap(ctx,
			item.Creator,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestApiCountMapRemove(t *testing.T) {
	keeper, ctx := keepertest.EstimatorKeeper(t)
	items := createNApiCountMap(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveApiCountMap(ctx,
			item.Creator,
		)
		_, found := keeper.GetApiCountMap(ctx,
			item.Creator,
		)
		require.False(t, found)
	}
}

func TestApiCountMapGetAll(t *testing.T) {
	keeper, ctx := keepertest.EstimatorKeeper(t)
	items := createNApiCountMap(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllApiCountMap(ctx)),
	)
}
