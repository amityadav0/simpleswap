package keeper

import (
	"encoding/binary"
	"fmt"

	"simpleswap/x/simpleswap/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (k Keeper) initializePool(ctx sdk.Context, msg *types.MsgCreatePool) (*string, error) {
	poolCreator := sdk.MustAccAddressFromBech32(msg.Creator)
	pool := msg.CreatePool()
	totalShares := sdk.NewInt(0)

	poolShareBaseDenom := pool.PoolId
	poolShareDisplayDenom := pool.PoolId

	assets := make(map[string]types.PoolAsset)
	for _, liquidity := range msg.Liquidity {
		assets[liquidity.Token.Denom] = liquidity
		totalShares = totalShares.Add(liquidity.Token.Amount)
	}

	// Check pool already created or not
	if _, found := k.GetPool(ctx, pool.PoolId); found {
		return nil, types.ErrAlreadyCreatedPool
	}

	// Check balance.
	for _, liquidity := range msg.Liquidity {
		balance := k.bankKeeper.GetBalance(ctx, poolCreator, liquidity.Token.Denom)
		if balance.Amount.LT(liquidity.Token.Amount) {
			return nil, types.ErrInsufficientBalance
		}
	}

	liquidity := sdk.NewCoins()
	for _, asset := range msg.Liquidity {
		liquidity = liquidity.Add(asset.Token)
	}
	liquidity.Sort()

	// Move asset from Sender to module account
	if err := k.LockTokens(ctx, pool.PoolId, poolCreator, liquidity); err != nil {
		return nil, err
	}

	// Register metadata to bank keeper
	k.bankKeeper.SetDenomMetaData(ctx, banktypes.Metadata{
		Description: fmt.Sprintf("The share token of the simple pool %s", pool.GetPoolId()),
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    poolShareBaseDenom,
				Exponent: 0,
				Aliases: []string{
					"poolshare",
				},
			},
			{
				Denom:    poolShareDisplayDenom,
				Exponent: types.OneShareExponent,
				Aliases:  nil,
			},
		},
		Base:    poolShareBaseDenom,
		Display: poolShareDisplayDenom,
	})

	// Mint shares
	if err := k.MintTokens(ctx, poolCreator, sdk.NewCoin(
		poolShareBaseDenom,
		totalShares,
	)); err != nil {
		return nil, err
	}

	// Save pool to chain
	k.AppendPool(ctx, pool)
	return &pool.PoolId, nil
}

func (k Keeper) GetPoolCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPoolsPrefix)
	b := store.Get(types.KeyCurrentPoolCountPrefix)
	if b == nil {
		return 0
	}
	return binary.BigEndian.Uint64(b)
}

func (k Keeper) SetPoolCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPoolsPrefix)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, count)
	store.Set(types.KeyCurrentPoolCountPrefix, b)
}

func GetPoolKey(count uint64) []byte {
	return []byte(fmt.Sprintf("%020d", count))
}

func (k Keeper) GetAllPool(ctx sdk.Context) (list []types.PoolData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(string(types.KeyPoolsPrefix)))

	// Start from the latest pool and move to the oldest
	poolCount := k.GetPoolCount(ctx)
	for i := poolCount; i >= 1 && (poolCount-i) < types.MaxPoolCount; i-- {
		b := store.Get(GetPoolKey(i))
		if b == nil {
			continue
		}
		var val types.PoolData
		k.cdc.MustUnmarshal(b, &val)
		list = append(list, val)
	}
	return
}

// Sets the mapping between poolId and its count index
func (k Keeper) SetPoolIDToCountMapping(ctx sdk.Context, poolID string, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPoolIDToCountPrefix)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, count)
	store.Set([]byte(poolID), b)
}

// Gets the count index of the poolId
func (k Keeper) GetCountByPoolID(ctx sdk.Context, poolID string) (count uint64, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPoolIDToCountPrefix)
	b := store.Get([]byte(poolID))
	if b == nil {
		return 0, false
	}
	return binary.BigEndian.Uint64(b), true
}

func (k Keeper) AppendPool(ctx sdk.Context, pool types.PoolData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPoolsPrefix)

	// Get current pool count
	poolCount := k.GetPoolCount(ctx)

	// Increment the count
	poolCount++

	// Set the new count
	k.SetPoolCount(ctx, poolCount)

	// Set the poolId to count mapping
	k.SetPoolIDToCountMapping(ctx, pool.PoolId, poolCount)

	// Marshal the pool and set in store
	b := k.cdc.MustMarshal(&pool)
	store.Set(GetPoolKey(poolCount), b)

	// Check if we exceed max pools
	if poolCount > types.MaxPoolCount {
		// Remove the oldest pool
		store.Delete(GetPoolKey(poolCount - types.MaxPoolCount))
	}
}

func (k Keeper) SetPool(ctx sdk.Context, pool types.PoolData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(string(types.KeyPoolsPrefix)))
	// Get current pool count
	poolCount, found := k.GetCountByPoolID(ctx, pool.PoolId)
	if !found {
		return
	}
	// Marshal the pool and set in store
	b := k.cdc.MustMarshal(&pool)
	store.Set(GetPoolKey(poolCount), b)
}

func (k Keeper) GetPool(ctx sdk.Context, poolID string) (val types.PoolData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPoolsPrefix)

	count, found := k.GetCountByPoolID(ctx, poolID)
	if !found {
		return val, false
	}

	b := store.Get(GetPoolKey(count))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
