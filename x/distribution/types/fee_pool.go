package types

import (
	"fmt"

	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
)

// zero fee pool
func InitialFeePool() FeePool {
	return FeePool{
		CommunityPool: sdk.DecCoins{},
	}
}

// ValidateGenesis validates the fee pool for a genesis state
func (f FeePool) ValidateGenesis() error {
	if f.CommunityPool.IsAnyNegative() {
		return fmt.Errorf("negative CommunityPool in distribution fee pool, is %v",
			f.CommunityPool)
	}

	return nil
}
