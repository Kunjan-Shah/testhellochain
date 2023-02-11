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

func createNApiHits(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ApiHits {
	items := make([]types.ApiHits, n)
	for i := range items {
		items[i].Id = keeper.AppendApiHits(ctx, items[i])
	}
	return items
}

func TestApiHitsGet(t *testing.T) {
	keeper, ctx := keepertest.EstimatorKeeper(t)
	items := createNApiHits(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetApiHits(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestApiHitsRemove(t *testing.T) {
	keeper, ctx := keepertest.EstimatorKeeper(t)
	items := createNApiHits(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveApiHits(ctx, item.Id)
		_, found := keeper.GetApiHits(ctx, item.Id)
		require.False(t, found)
	}
}

func TestApiHitsGetAll(t *testing.T) {
	keeper, ctx := keepertest.EstimatorKeeper(t)
	items := createNApiHits(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllApiHits(ctx)),
	)
}

func TestApiHitsCount(t *testing.T) {
	keeper, ctx := keepertest.EstimatorKeeper(t)
	items := createNApiHits(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetApiHitsCount(ctx))
}
