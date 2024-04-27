package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetEventAttrOfAsset(assets []sdk.Coin) []sdk.Attribute {
	attr := []sdk.Attribute{}
	for index, asset := range assets {
		attr = append(attr, sdk.NewAttribute(
			fmt.Sprintf("%d", index),
			asset.String(),
		))
	}
	return attr
}

func GetLiquidityAsCoins(tokens []sdk.Coin) sdk.Coins {
	return sdk.NewCoins(tokens...).Sort()
}

const (
	EventValueActionCreatePool   = "create_pool"
	EventValueActionAddLiquidity = "add_liquidity"
	EventValueActionWithdraw     = "withdraw"
	EventValueActionSwap         = "swap"
)

const (
	AttributeKeyAction      = "action"
	AttributeKeyPoolID      = "pool_id"
	AttributeKeyTokenIn     = "token_in"
	AttributeKeyTokenOut    = "token_out"
	AttributeKeyLpToken     = "liquidity_pool_token"
	AttributeKeyLpSupply    = "liquidity_pool_token_supply"
	AttributeKeyPoolCreator = "pool_creator"
	AttributeKeyName        = "name"
	AttributeKeyMsgSender   = "msg_sender"
)
