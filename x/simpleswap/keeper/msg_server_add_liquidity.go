package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"simpleswap/x/simpleswap/types"
)

func (k msgServer) AddLiquidity(goCtx context.Context, msg *types.MsgAddLiquidity) (*types.MsgAddLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddLiquidityResponse{}, nil
}
