package keeper

import (
	"testhellochain/x/estimator/types"
)

var _ types.QueryServer = Keeper{}
