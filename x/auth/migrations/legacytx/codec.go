package legacytx

import (
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(StdTx{}, "cosmos-sdk/StdTx", nil)
}
