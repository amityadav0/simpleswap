package types

import (
	"crypto/sha256"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (p *PoolData) GetAssetDenoms() []string {
	denoms := []string{}
	for _, asset := range p.Assets {
		denoms = append(denoms, asset.Token.Denom)
	}
	return denoms
}

func GetEscrowAddress(poolID string) sdk.AccAddress {
	// a slash is used to create domain separation between port and channel identifiers to
	// prevent address collisions between escrow addresses

	// ADR 028 AddressHash construction
	preImage := []byte("1")
	preImage = append(preImage, 0)
	preImage = append(preImage, poolID...)
	hash := sha256.Sum256(preImage)
	return hash[:20]
}

// Helper functions
func (p *PoolData) TakeFees(amount sdkmath.Int) sdk.Dec {
	amountDec := sdk.NewDecFromInt(amount)
	feeRate := p.PoolParams.SwapFee.Quo(sdk.NewDec(10000))
	fees := amountDec.Mul(feeRate)
	amountMinusFees := amountDec.Sub(fees)
	return amountMinusFees
}

// IncreaseShare add xx amount share to total share amount in pool
func (p *PoolData) IncreaseShare(amt sdkmath.Int) {
	p.TotalShares.Amount = p.TotalShares.Amount.Add(amt)
}

// DecreaseShare subtract xx amount share to total share amount in pool
func (p *PoolData) DecreaseShare(amt sdkmath.Int) {
	p.TotalShares.Amount = p.TotalShares.Amount.Sub(amt)
}

// IncreaseLiquidity adds xx amount liquidity to assets in pool
func (p *PoolData) IncreaseLiquidity(coins []sdk.Coin) error {
	for _, coin := range coins {
		asset, exists := p.Assets[coin.Denom]
		if !exists {
			return ErrNotFoundAssetInPool
		}
		// Add liquidity logic here
		asset.Token.Amount = asset.Token.Amount.Add(coin.Amount)
		p.Assets[coin.Denom] = asset
	}
	// Update TotalShares or other fields if necessary
	return nil
}

// DecreaseLiquidity subtracts xx amount liquidity from assets in pool
func (p *PoolData) DecreaseLiquidity(coins []sdk.Coin) error {
	for _, coin := range coins {
		asset, exists := p.Assets[coin.Denom]
		if !exists {
			return ErrNotFoundAssetInPool
		}
		// Add liquidity logic here
		asset.Token.Amount = asset.Token.Amount.Sub(coin.Amount)
		p.Assets[coin.Denom] = asset
	}
	// Update TotalShares or other fields if necessary
	return nil
}

func (p *PoolData) GetAssetList() []PoolAsset {
	assets := make([]PoolAsset, 0)
	if p != nil {
		for _, asset := range p.Assets {
			assets = append(assets, asset)
		}
		return assets
	}
	return nil
}

func (p *PoolData) GetTokens() []sdk.Coin {
	assets := make([]sdk.Coin, 0)
	if p != nil {
		for _, asset := range p.Assets {
			assets = append(assets, asset.Token)
		}
		return assets
	}
	return nil
}

func (p *PoolData) GetLiquidity() sdk.Coins {
	assets := sdk.NewCoins()
	if p != nil {
		for _, asset := range p.Assets {
			assets = assets.Add(asset.Token)
		}
		return assets
	}
	return nil
}

func (p *PoolData) Sum() sdkmath.Int {
	sum := sdkmath.ZeroInt()
	if p != nil {
		for _, asset := range p.Assets {
			sum = sum.Add(asset.Token.Amount)
		}
		return sum
	}
	return sdk.ZeroInt()
}
