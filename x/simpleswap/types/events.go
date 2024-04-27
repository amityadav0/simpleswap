package types

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
