package keeper

import (
	"context"

	"simpleswap/x/simpleswap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddLiquidity(goCtx context.Context, msg *types.MsgAddLiquidity) (*types.MsgAddLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	pool, found := k.GetPool(ctx, msg.PoolId)
	if !found {
		return nil, types.ErrNotFoundAssetInPool
	}
	sender := sdk.MustAccAddressFromBech32(msg.Creator)

	// Move asset from sender to module account

	if err := k.LockTokens(ctx, msg.PoolId, sender, types.GetLiquidityAsCoins(msg.Liquidity)); err != nil {
		return nil, err
	}

	// Mint share to sender
	share, err := pool.EstimateShare(msg.Liquidity)
	if err != nil {
		return nil, err
	}
	if err := k.MintTokens(ctx, sender, share); err != nil {
		return nil, err
	}
	// Update pool status
	if err := pool.IncreaseLiquidity(msg.Liquidity); err != nil {
		return nil, err
	}
	pool.IncreaseShare(share.Amount)

	// Save update information
	k.SetPool(ctx, pool)

	// Emit events
	k.EmitEvent(
		ctx, types.EventValueActionAddLiquidity,
		msg.PoolId,
		msg.Creator,
		types.GetEventAttrOfAsset(msg.Liquidity)...,
	)
	return &types.MsgAddLiquidityResponse{}, nil
}
