package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"simpleswap/x/simpleswap/types"
)

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSwapResponse{}, nil
}
