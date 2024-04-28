package keeper_test

import (
	simapp "simpleswap/app"
	"simpleswap/testutil/sample"
	"simpleswap/x/simpleswap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestMsgAddLiquidity() {
	suite.SetupTest()

	tests := []struct {
		name    string
		mutator func(*types.MsgAddLiquidity, string)
	}{
		{
			"add liquidity to pool",
			func(msg *types.MsgAddLiquidity, poolID string) {
				msg.Liquidity = []sdk.Coin{
					sdk.NewCoin(simapp.WETH, sdk.NewInt(100)),
					sdk.NewCoin(simapp.ETH, sdk.NewInt(100)),
				}
			},
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			// Create a new pool of the specific type
			poolID := suite.CreateNewPool()

			// Initialize the MsgAddLiquidity
			msg := types.MsgAddLiquidity{
				Creator: types.Alice,
				PoolId:  poolID,
			}

			// Use the mutator to set the liquidity for the specific pool type
			tc.mutator(&msg, poolID)

			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.msgServer.AddLiquidity(ctx, &msg)

			suite.Require().NoError(err)
			suite.Require().NotNil(res)
		})
	}
}

func (suite *KeeperTestSuite) TestMsgAddLiquidityFail() {
	var msg *types.MsgAddLiquidity
	// Create a new pool
	poolID := suite.CreateNewPool()

	testCases := []struct {
		name   string
		mallet func(msg *types.MsgAddLiquidity, poolID string)
	}{
		{
			"invalid sender",
			func(msg *types.MsgAddLiquidity, poolID string) {},
		},
		{
			"invalid poolID",
			func(msg *types.MsgAddLiquidity, poolID string) {
				msg = &types.MsgAddLiquidity{
					Creator: sample.AccAddress(),
					PoolId:  "",
					Liquidity: []sdk.Coin{
						sdk.NewCoin(
							simapp.ETH,
							sdk.NewInt(100),
						),
						sdk.NewCoin(
							simapp.WETH,
							sdk.NewInt(100),
						),
					},
				}
			},
		},
		{
			"not enough funds",
			func(msg *types.MsgAddLiquidity, poolID string) {
				msg = &types.MsgAddLiquidity{
					Creator: sample.AccAddress(),
					PoolId:  poolID,
					Liquidity: []sdk.Coin{
						sdk.NewCoin(
							simapp.ETH,
							sdk.NewInt(100),
						),
						sdk.NewCoin(
							simapp.WETH,
							sdk.NewInt(100),
						),
					},
				}
			},
		},
		{
			"invalid asset type",
			func(msg *types.MsgAddLiquidity, poolID string) {
				msg.Liquidity = []sdk.Coin{
					sdk.NewCoin("INVALID_ASSET_TYPE", sdk.NewInt(100)),
				}
			},
		},
		{
			"zero liquidity",
			func(msg *types.MsgAddLiquidity, poolID string) {
				msg.Liquidity = []sdk.Coin{
					sdk.NewCoin(simapp.DefaultBondDenom, sdk.NewInt(0)),
				}
			},
		},
	}

	for _, tc := range testCases {

		msg = types.NewMsgAddLiquidity(
			"",
			poolID,
			[]sdk.Coin{},
		)
		tc.mallet(msg, poolID)

		res, err := suite.msgServer.AddLiquidity(sdk.WrapSDKContext(suite.ctx), msg)
		suite.Require().Error(err)
		suite.Require().Nil(res)
	}
}
