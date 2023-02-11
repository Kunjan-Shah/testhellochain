package testhellochain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "testhellochain/testutil/keeper"
	"testhellochain/testutil/nullify"
	"testhellochain/x/testhellochain"
	"testhellochain/x/testhellochain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TesthellochainKeeper(t)
	testhellochain.InitGenesis(ctx, *k, genesisState)
	got := testhellochain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
