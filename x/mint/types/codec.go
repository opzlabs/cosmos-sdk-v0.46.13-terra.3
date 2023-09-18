package types

import (
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec"
	cryptocodec "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/crypto/codec"
)

var amino = codec.NewLegacyAmino()

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
