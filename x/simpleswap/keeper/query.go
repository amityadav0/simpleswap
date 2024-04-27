package keeper

import (
	"simpleswap/x/simpleswap/types"
)

var _ types.QueryServer = Keeper{}
