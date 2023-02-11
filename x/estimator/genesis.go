package estimator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testhellochain/x/estimator/keeper"
	"testhellochain/x/estimator/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the apiHits
	for _, elem := range genState.ApiHitsList {
		k.SetApiHits(ctx, elem)
	}

	// Set apiHits count
	k.SetApiHitsCount(ctx, genState.ApiHitsCount)
	// Set all the apiData
	for _, elem := range genState.ApiDataList {
		k.SetApiData(ctx, elem)
	}

	// Set apiData count
	k.SetApiDataCount(ctx, genState.ApiDataCount)
	// Set all the apiCountMap
	for _, elem := range genState.ApiCountMapList {
		k.SetApiCountMap(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ApiHitsList = k.GetAllApiHits(ctx)
	genesis.ApiHitsCount = k.GetApiHitsCount(ctx)
	genesis.ApiDataList = k.GetAllApiData(ctx)
	genesis.ApiDataCount = k.GetApiDataCount(ctx)
	genesis.ApiCountMapList = k.GetAllApiCountMap(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
