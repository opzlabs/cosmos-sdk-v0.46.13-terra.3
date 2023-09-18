package types

import (
	sdkerrors "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types/errors"
)

// x/crisis module sentinel errors
var (
	ErrNoSender         = sdkerrors.Register(ModuleName, 2, "sender address is empty")
	ErrUnknownInvariant = sdkerrors.Register(ModuleName, 3, "unknown invariant")
)
