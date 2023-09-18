package types

import (
	"fmt"

	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	paramtypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/params/types"
)

// ParamStoreKeyConstantFee is the key for constant fee parameter
var ParamStoreKeyConstantFee = []byte("ConstantFee")

// type declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable(
		paramtypes.NewParamSetPair(ParamStoreKeyConstantFee, sdk.Coin{}, validateConstantFee),
	)
}

func validateConstantFee(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() {
		return fmt.Errorf("invalid constant fee: %s", v)
	}

	return nil
}
