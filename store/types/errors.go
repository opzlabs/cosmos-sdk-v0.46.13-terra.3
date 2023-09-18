package types

import (
	sdkerrors "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types/errors"
)

const StoreCodespace = "store"

var ErrInvalidProof = sdkerrors.Register(StoreCodespace, 2, "invalid proof")
