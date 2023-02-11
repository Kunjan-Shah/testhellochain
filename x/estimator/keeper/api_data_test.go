package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "testhellochain/testutil/keeper"
	"testhellochain/testutil/nullify"
	"testhellochain/x/estimator/keeper"
	"testhellochain/x/estimator/types"
)

func createNApiData(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ApiData {
	items := make([]types.ApiData, n)
	for i := range items {
		items[i].Id = keeper.AppendApiData(ctx, items[i])
	}
	return items
}

func TestApiDataGet(t *testing.T) {
	keeper, ctx := keepertest.EstimatorKeeper(t)
	items := createNApiData(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetApiData(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestApiDataRemove(t *testing.T) {
	keeper, ctx := keepertest.EstimatorKeeper(t)
	items := createNApiData(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveApiData(ctx, item.Id)
		_, found := keeper.GetApiData(ctx, item.Id)
		require.False(t, found)
	}
}

func TestApiDataGetAll(t *testing.T) {
	keeper, ctx := keepertest.EstimatorKeeper(t)
	items := createNApiData(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllApiData(ctx)),
	)
}

func TestApiDataCount(t *testing.T) {
	keeper, ctx := keepertest.EstimatorKeeper(t)
	items := createNApiData(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetApiDataCount(ctx))
}
