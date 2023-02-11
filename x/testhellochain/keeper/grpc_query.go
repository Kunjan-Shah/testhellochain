package keeper

import (
	"testhellochain/x/testhellochain/types"
)

var _ types.QueryServer = Keeper{}
