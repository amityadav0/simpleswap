package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "simpleswap/testutil/keeper"
	"simpleswap/x/simpleswap/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SimpleswapKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
