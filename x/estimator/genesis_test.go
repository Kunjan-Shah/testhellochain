package estimator_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "testhellochain/testutil/keeper"
	"testhellochain/testutil/nullify"
	"testhellochain/x/estimator"
	"testhellochain/x/estimator/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ApiHitsList: []types.ApiHits{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ApiHitsCount: 2,
		ApiDataList: []types.ApiData{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ApiDataCount: 2,
		ApiCountMapList: []types.ApiCountMap{
			{
				Creator: "0",
			},
			{
				Creator: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EstimatorKeeper(t)
	estimator.InitGenesis(ctx, *k, genesisState)
	got := estimator.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ApiHitsList, got.ApiHitsList)
	require.Equal(t, genesisState.ApiHitsCount, got.ApiHitsCount)
	require.ElementsMatch(t, genesisState.ApiDataList, got.ApiDataList)
	require.Equal(t, genesisState.ApiDataCount, got.ApiDataCount)
	require.ElementsMatch(t, genesisState.ApiCountMapList, got.ApiCountMapList)
	// this line is used by starport scaffolding # genesis/test/assert
}
