package keyring

import (
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec/legacy"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/crypto/hd"
)

func init() {
	RegisterLegacyAminoCodec(legacy.Cdc)
}

// RegisterLegacyAminoCodec registers concrete types and interfaces on the given codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*LegacyInfo)(nil), nil)
	cdc.RegisterConcrete(hd.BIP44Params{}, "crypto/keys/hd/BIP44Params", nil)
	cdc.RegisterConcrete(legacyLocalInfo{}, "crypto/keys/localInfo", nil)
	cdc.RegisterConcrete(legacyLedgerInfo{}, "crypto/keys/ledgerInfo", nil)
	cdc.RegisterConcrete(legacyOfflineInfo{}, "crypto/keys/offlineInfo", nil)
	cdc.RegisterConcrete(LegacyMultiInfo{}, "crypto/keys/multiInfo", nil)
}
