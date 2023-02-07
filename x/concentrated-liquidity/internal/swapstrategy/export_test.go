package swapstrategy

import sdk "github.com/cosmos/cosmos-sdk/types"

func ComputeFeeChargePerSwapStepOutGivenIn(currentSqrtPrice sdk.Dec, hasReachedTarget bool, amountIn, amountSpecifiedRemaining, swapFee sdk.Dec) sdk.Dec {
	return computeFeeChargePerSwapStepOutGivenIn(currentSqrtPrice, hasReachedTarget, amountIn, amountSpecifiedRemaining, swapFee)
}
