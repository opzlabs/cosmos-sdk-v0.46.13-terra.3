package keeper_test

import (
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/simapp"
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	authtypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/auth/types"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/distribution/types"
)

var (
	PKS = simapp.CreateTestPubKeys(5)

	valConsPk1 = PKS[0]
	valConsPk2 = PKS[1]
	valConsPk3 = PKS[2]

	valConsAddr1 = sdk.ConsAddress(valConsPk1.Address())
	valConsAddr2 = sdk.ConsAddress(valConsPk2.Address())

	distrAcc = authtypes.NewEmptyModuleAccount(types.ModuleName)
)
