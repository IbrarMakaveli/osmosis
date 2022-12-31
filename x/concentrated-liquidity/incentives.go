package concentrated_liquidity

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/osmoutils/sumtree"
	"github.com/osmosis-labs/osmosis/v13/x/concentrated-liquidity/model"
	"github.com/osmosis-labs/osmosis/v13/x/concentrated-liquidity/types"
)

// Gets the current join-time sumtree for the passed in pool ID.
func (k Keeper) getSumtree(ctx sdk.Context, poolID uint64) sumtree.Tree {
	return sumtree.NewTree(prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyAccumulationStore(poolID)), 10)
}

// Adds to a position's liquidity in Liquidity tree, setting its JoinTime to current block time.
//
// Since we want to reset a position's JoinTime every time they update their liquidity, any
// liquidity update is equivalent to removing the old JoinTime node from the liquidity tree
// and replacing it with a new one. Thus, this function simply takes in the old and new liquidity
// amounts and leaves liquidity delta calculations to the caller.
func (k Keeper) addToLiquidityTree(ctx sdk.Context, poolId uint64, position *model.Position, oldLiquidity sdk.Dec, newLiquidity sdk.Dec, posOwner sdk.AccAddress, posLowerTick int64, posUpperTick int64) error {
	if newLiquidity.LT(oldLiquidity) {
		return fmt.Errorf("Attempted to add a negative amount of liquidity. Should be using `removeFromLiquidityTree`")
	}

	// Clear old liquidity from position tree
	// Note: we use Decrease instead of Remove to accommodate multiple positions with the same join time
	k.getSumtree(ctx, poolId).Decrease(types.KeyJoinTime(position.JoinTime), oldLiquidity)

	// Update position's JoinTime to current block time
	position.JoinTime = ctx.BlockTime()
	k.setPosition(ctx, poolId, posOwner, posLowerTick, posUpperTick, position)

	// Add new position time to JoinTime sumtree
	k.getSumtree(ctx, poolId).Increase(types.KeyJoinTime(position.JoinTime), newLiquidity)

	return nil
}

// Remove liquidity from `position`s join time without resetting position.JoinTime.
// This is to allow for LPs to remove a portion of their liquidity without forfeiting their uptime
// on their remaining liquidity.
func (k Keeper) removeFromLiquidityTree(ctx sdk.Context, poolId uint64, position *model.Position, liquidityToRemove sdk.Dec, posOwner sdk.AccAddress, posLowerTick int64, posUpperTick int64) error {
	if liquidityToRemove.GT(position.Liquidity) {
		return fmt.Errorf("Attempted to remove more liquidity than the position has")
	}

	if liquidityToRemove.IsPositive() {
		return fmt.Errorf("Liquidity delta for removing liquidity must be negative")
	}

	liquidityToRemove = liquidityToRemove.Neg()

	// Remove liquidity from position's JoinTime node in tree
	// Note: we use Decrease instead of Remove to accommodate multiple positions with the same join time
	k.getSumtree(ctx, poolId).Decrease(types.KeyJoinTime(position.JoinTime), liquidityToRemove)

	return nil
}

// Gets total liquidity that has joined at time <= `joinTime`
// nolint:unused
func (k Keeper) getLiquidityBeforeOrAtJoinTime(ctx sdk.Context, poolId uint64, joinTime time.Time) sdk.Dec {
	return k.getSumtree(ctx, poolId).PrefixSum(types.KeyJoinTime(joinTime))
}

// Gets total liquidity that joined after time `joinTime`
// nolint:unused
func (k Keeper) getLiquidityAfterJoinTime(ctx sdk.Context, poolId uint64, joinTime time.Time) sdk.Dec {
	return k.getSumtree(ctx, poolId).SubsetAccumulation(types.KeyJoinTime(joinTime.Add(1*time.Second)), nil)
}

// Gets all liquidity that joined exactly at `joinTime`. This is primarily intended for use in validation logic.
// nolint:unused
func (k Keeper) getLiquidityExactlyAtJoinTime(ctx sdk.Context, poolId uint64, joinTime time.Time) sdk.Dec {
	return k.getSumtree(ctx, poolId).SubsetAccumulation(types.KeyJoinTime(joinTime), types.KeyJoinTime(joinTime))
}
