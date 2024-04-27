package types

import (
	"crypto/sha256"
	"encoding/hex"
	fmt "fmt"
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreatePool = "create_pool"

var _ sdk.Msg = &MsgCreatePool{}

func NewMsgCreatePool(creator string, poolParams PoolParams, poolAsset []PoolAsset) *MsgCreatePool {
	return &MsgCreatePool{
		Creator:   creator,
		Params:    &poolParams,
		Liquidity: poolAsset,
	}
}

func (msg *MsgCreatePool) Route() string {
	return RouterKey
}

func (msg *MsgCreatePool) Type() string {
	return TypeMsgCreatePool
}

func (msg *MsgCreatePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgCreatePool) CreatePool() PoolData {
	// Extract denom list from Liquidity
	denoms := msg.GetAssetDenoms()

	assets := make(map[string]PoolAsset)
	totalShares := sdk.NewInt(0)
	for _, liquidity := range msg.Liquidity {
		assets[liquidity.Token.Denom] = liquidity
		totalShares = totalShares.Add(liquidity.Token.Amount)
	}

	// Generate poolID
	sort.Strings(denoms)
	poolIDHash := sha256.New()
	poolIDHash.Write([]byte(strings.Join(denoms, "")))
	newPoolID := "pool" + fmt.Sprintf("%v", hex.EncodeToString(poolIDHash.Sum(nil)))

	poolShareBaseDenom := GetPoolShareDenom(newPoolID)
	pool := PoolData{
		PoolId:      newPoolID,
		PoolParams:  *msg.Params,
		Assets:      assets,
		TotalShares: sdk.NewCoin(poolShareBaseDenom, totalShares),
	}
	return pool
}

func (msg *MsgCreatePool) GetAssetDenoms() []string {
	denoms := []string{}
	for _, asset := range msg.Liquidity {
		denoms = append(denoms, asset.Token.Denom)
	}
	return denoms
}
