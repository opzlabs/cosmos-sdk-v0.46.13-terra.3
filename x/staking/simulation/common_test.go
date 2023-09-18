package simulation_test

import (
	"math/big"

	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
)

func init() {
	sdk.DefaultPowerReduction = sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
}
