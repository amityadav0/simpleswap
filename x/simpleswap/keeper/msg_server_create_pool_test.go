package keeper_test

import (
	simapp "simpleswap/app"
	"simpleswap/testutil/sample"
	"simpleswap/x/simpleswap/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestMsgCreatePool() {
	suite.SetupTest()

	tests := []struct {
		name    string
		mutator func(msg *types.MsgCreatePool)
	}{
		{
			"pool",
			func(msg *types.MsgCreatePool) {
				amp := sdk.NewInt(100)
				msg.Params.Amp = &amp
				msg.Liquidity = suite.createPoolLiquidity()
			},
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			msg := suite.defaultMsgCreatePool()
			tc.mutator(msg)

			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.msgServer.CreatePool(ctx, msg)

			suite.Require().NoError(err)
			suite.Require().NotNil(res)

			// Check pool is created or not
			pool, err := suite.queryClient.Pools(ctx, &types.QueryPoolsRequest{})
			suite.Require().NoError(err)
			suite.Require().Equal(pool.Pools[0].PoolId, res.PoolId)
		})
	}
}

// Helper method to create a default MsgCreatePool for testing
func (suite *KeeperTestSuite) defaultMsgCreatePool() *types.MsgCreatePool {
	amp := sdk.NewInt(100)
	return types.NewMsgCreatePool(
		types.Alice,
		types.PoolParams{
			SwapFee: sdkmath.LegacyDec(sdk.NewInt(100)),
			Amp:     &amp,
		},
		[]types.PoolAsset{
			{
				Token:   sdk.NewCoin(simapp.WETH, sdkmath.NewInt(100)),
				Decimal: sdk.NewInt(6),
			},
			{
				Token:   sdk.NewCoin(simapp.ETH, sdkmath.NewInt(100)),
				Decimal: sdk.NewInt(6),
			},
		},
	)
}

// Helper method to create pool liquidity for testing
func (suite *KeeperTestSuite) createPoolLiquidity() []types.PoolAsset {
	return []types.PoolAsset{
		{
			Token:   sdk.NewCoin(simapp.WETH, sdkmath.NewInt(1000)),
			Decimal: sdk.NewInt(6),
		},
		{
			Token:   sdk.NewCoin(simapp.ETH, sdkmath.NewInt(1000)),
			Decimal: sdk.NewInt(6),
		},
	}
}

func (suite *KeeperTestSuite) TestMsgCreatePoolFail() {
	var msg *types.MsgCreatePool
	suite.SetupTest()
	amp := sdk.NewInt(100)

	testCases := []struct {
		name   string
		mallet func()
	}{
		{
			"invalid sender",
			func() {},
		},
		{
			"not enough funds",
			func() {
				msg = types.NewMsgCreatePool(
					sample.AccAddress(),
					types.PoolParams{
						SwapFee: sdkmath.LegacyDec(sdk.NewInt(100)),
						Amp:     &amp,
					},
					[]types.PoolAsset{
						{
							Token:   sdk.NewCoin(simapp.WETH, sdkmath.NewInt(100)),
							Decimal: sdk.NewInt(6),
						},
						{
							Token:   sdk.NewCoin(simapp.ETH, sdkmath.NewInt(100)),
							Decimal: sdk.NewInt(6),
						},
					},
				)
			},
		},
	}

	for _, tc := range testCases {

		msg = types.NewMsgCreatePool(
			"",
			types.PoolParams{},
			[]types.PoolAsset{},
		)
		tc.mallet()

		res, err := suite.msgServer.CreatePool(sdk.WrapSDKContext(suite.ctx), msg)
		suite.Require().Error(err)
		suite.Require().Nil(res)
	}
}

func (suite *KeeperTestSuite) CreateNewPool() string {
	var msg *types.MsgCreatePool
	suite.SetupTest()

	amp := sdk.NewInt(1)

	msg = types.NewMsgCreatePool(
		types.Alice,
		types.PoolParams{
			SwapFee: sdkmath.LegacyDec(sdk.NewInt(100)),
			Amp:     &amp,
		},
		[]types.PoolAsset{
			{
				Token:   sdk.NewCoin(simapp.WETH, sdkmath.NewInt(1000)),
				Decimal: sdk.NewInt(6),
			},
			{
				Token:   sdk.NewCoin(simapp.ETH, sdkmath.NewInt(1000)),
				Decimal: sdk.NewInt(6),
			},
		},
	)

	ctx := sdk.WrapSDKContext(suite.ctx)
	res, err := suite.msgServer.CreatePool(ctx, msg)

	suite.Require().NoError(err)
	suite.Require().NotNil(res)
	return res.PoolId
}
