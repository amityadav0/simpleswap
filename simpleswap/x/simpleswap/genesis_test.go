package simpleswap_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "simpleswap/testutil/keeper"
	"simpleswap/testutil/nullify"
	"simpleswap/x/simpleswap"
	"simpleswap/x/simpleswap/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SimpleswapKeeper(t)
	simpleswap.InitGenesis(ctx, *k, genesisState)
	got := simpleswap.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
