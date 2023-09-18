package keys

import (
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec"
	cryptocodec "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/crypto/codec"
)

// TODO: remove this file https://github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/issues/8047

// KeysCdc defines codec to be used with key operations
var KeysCdc *codec.LegacyAmino

func init() {
	KeysCdc = codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(KeysCdc)
	KeysCdc.Seal()
}

// marshal keys
func MarshalJSON(o interface{}) ([]byte, error) {
	return KeysCdc.MarshalJSON(o)
}

// unmarshal json
func UnmarshalJSON(bz []byte, ptr interface{}) error {
	return KeysCdc.UnmarshalJSON(bz, ptr)
}
