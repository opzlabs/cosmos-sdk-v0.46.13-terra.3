package codec

import (
	"github.com/opzlabs/cosmos-sdk/codec"
	cryptocodec "github.com/opzlabs/cosmos-sdk/crypto/codec"
	sdk "github.com/opzlabs/cosmos-sdk/types"
)

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	cryptocodec.RegisterCrypto(Amino)
	codec.RegisterEvidences(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
}
