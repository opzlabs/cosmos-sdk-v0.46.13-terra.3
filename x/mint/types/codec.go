package types

import (
	"github.com/opzlabs/cosmos-sdk/codec"
	cryptocodec "github.com/opzlabs/cosmos-sdk/crypto/codec"
)

var amino = codec.NewLegacyAmino()

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
