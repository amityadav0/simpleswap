package keeper

import (
	"context"
	"encoding/binary"

	"simpleswap/x/simpleswap/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Pools(goCtx context.Context, req *types.QueryPoolsRequest) (*types.QueryPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	pools := []types.PoolData{}
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Use the store with the mapping of poolId to its count
	store := ctx.KVStore(k.storeKey)
	poolIDToCountStore := prefix.NewStore(store, types.KeyPoolIDToCountPrefix)

	// A counter for pagination
	var counter int

	pageRes, err := query.Paginate(poolIDToCountStore, req.Pagination, func(key []byte, value []byte) error {
		counter++

		// Decode the count for the given poolId
		count := binary.BigEndian.Uint64(value)

		// Get the actual pool using the count
		poolStore := prefix.NewStore(store, types.KeyPrefix(string(types.KeyPoolsPrefix)))
		poolBytes := poolStore.Get(GetPoolKey(count))
		if poolBytes == nil {
			return nil
		}

		var pool types.PoolData
		if err := k.cdc.Unmarshal(poolBytes, &pool); err != nil {
			return nil
		}
		pools = append(pools, pool)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPoolsResponse{Pools: pools, Pagination: pageRes}, nil
}
