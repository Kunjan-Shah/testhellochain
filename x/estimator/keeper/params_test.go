package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "testhellochain/testutil/keeper"
	"testhellochain/x/estimator/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EstimatorKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
