package keeper

import (
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	authtypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/auth/types"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/distribution/types"
)

// get outstanding rewards
func (k Keeper) GetValidatorOutstandingRewardsCoins(ctx sdk.Context, val sdk.ValAddress) sdk.DecCoins {
	return k.GetValidatorOutstandingRewards(ctx, val).Rewards
}

// get the community coins
func (k Keeper) GetFeePoolCommunityCoins(ctx sdk.Context) sdk.DecCoins {
	return k.GetFeePool(ctx).CommunityPool
}

// GetDistributionAccount returns the distribution ModuleAccount
func (k Keeper) GetDistributionAccount(ctx sdk.Context) authtypes.ModuleAccountI {
	return k.authKeeper.GetModuleAccount(ctx, types.ModuleName)
}
