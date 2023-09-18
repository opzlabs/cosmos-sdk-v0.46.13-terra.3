package v046_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/simapp"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/testutil"
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	paramtypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/params/types"
	v046staking "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/staking/migrations/v046"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/staking/types"
)

func TestStoreMigration(t *testing.T) {
	encCfg := simapp.MakeTestEncodingConfig()
	stakingKey := sdk.NewKVStoreKey("staking")
	tStakingKey := sdk.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(stakingKey, tStakingKey)
	paramstore := paramtypes.NewSubspace(encCfg.Codec, encCfg.Amino, stakingKey, tStakingKey, "staking")

	// Check no params
	require.False(t, paramstore.Has(ctx, types.KeyMinCommissionRate))

	// Run migrations.
	err := v046staking.MigrateStore(ctx, stakingKey, encCfg.Codec, paramstore)
	require.NoError(t, err)

	// Make sure the new params are set.
	require.True(t, paramstore.Has(ctx, types.KeyMinCommissionRate))
}
