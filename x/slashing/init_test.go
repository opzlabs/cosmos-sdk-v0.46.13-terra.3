package slashing_test

import (
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
)

// The default power validators are initialized to have within tests
var InitTokens = sdk.TokensFromConsensusPower(200, sdk.DefaultPowerReduction)
