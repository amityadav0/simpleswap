package keeper

import (
	"fmt"

	"simpleswap/x/simpleswap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) LockTokens(ctx sdk.Context, poolID string, sender sdk.AccAddress, tokens sdk.Coins) error {
	escrow := types.GetEscrowAddress(poolID)
	return k.bankKeeper.SendCoins(
		ctx, sender, escrow, tokens,
	)
}

func (k Keeper) UnLockTokens(ctx sdk.Context, poolID string, receiver sdk.AccAddress, tokens sdk.Coins) error {
	escrow := types.GetEscrowAddress(poolID)
	return k.bankKeeper.SendCoins(
		ctx, escrow, receiver, sdk.NewCoins(tokens...),
	)
}

func (k Keeper) BurnTokens(ctx sdk.Context, sender sdk.AccAddress, tokens sdk.Coin) error {
	// transfer coins to module account and burn them
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(tokens)); err != nil {
		return err
	}
	if err := k.bankKeeper.BurnCoins(
		ctx, types.ModuleName, sdk.NewCoins(tokens),
	); err != nil {
		// NOTE: should not happen as the module account was
		// retrieved on the step above and it has enough balance
		// to burn.
		panic(fmt.Sprintf("cannot burn coins after a successful send to a module account: %v", err))
	}
	return nil
}

func (k Keeper) MintTokens(ctx sdk.Context, receiver sdk.AccAddress, tokens sdk.Coin) error {
	// mint new tokens if the source of the transfer is the same chain
	if err := k.bankKeeper.MintCoins(
		ctx, types.ModuleName, sdk.NewCoins(tokens),
	); err != nil {
		return err
	}
	// send to receiver
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, receiver, sdk.NewCoins(tokens),
	); err != nil {
		panic(fmt.Sprintf("unable to send coins from module to account: %v", err))
	}
	return nil
}
