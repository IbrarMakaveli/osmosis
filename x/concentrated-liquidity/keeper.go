package concentrated_liquidity

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/osmosis-labs/osmosis/v13/x/concentrated-liquidity/types"
	swaproutertypes "github.com/osmosis-labs/osmosis/v13/x/swaprouter/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryCodec

	paramSpace paramtypes.Subspace

	// keepers
	bankKeeper types.BankKeeper
}

func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey, bankKeeper types.BankKeeper, paramSpace paramtypes.Subspace) *Keeper {
	// ParamSubspace must be initialized within app/keepers/keepers.go
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}
	return &Keeper{
		storeKey:   storeKey,
		paramSpace: paramSpace,
		cdc:        cdc,
		bankKeeper: bankKeeper,
	}
}

// GetParams returns the total set of concentrated-liquidity module's parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the concentrated-liquidity module's parameters with the provided parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// TODO: spec, tests, implementation
func (k Keeper) InitializePool(ctx sdk.Context, pool swaproutertypes.PoolI, creatorAddress sdk.AccAddress) error {
	panic("not implemented")
}