package ledger

import (
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec"
	cryptoAmino "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/crypto/codec"
)

var cdc = codec.NewLegacyAmino()

func init() {
	RegisterAmino(cdc)
	cryptoAmino.RegisterCrypto(cdc)
}

// RegisterAmino registers all go-crypto related types in the given (amino) codec.
func RegisterAmino(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(PrivKeyLedgerSecp256k1{},
		"tendermint/PrivKeyLedgerSecp256k1", nil)
}
