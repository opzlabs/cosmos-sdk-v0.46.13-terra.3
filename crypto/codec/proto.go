package codec

import (
	codectypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec/types"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/crypto/keys/ed25519"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/crypto/keys/multisig"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/crypto/keys/secp256k1"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/crypto/keys/secp256r1"
	cryptotypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/crypto/types"
)

// RegisterInterfaces registers the sdk.Tx interface.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	var pk *cryptotypes.PubKey
	registry.RegisterInterface("cosmos.crypto.PubKey", pk)
	registry.RegisterImplementations(pk, &ed25519.PubKey{})
	registry.RegisterImplementations(pk, &secp256k1.PubKey{})
	registry.RegisterImplementations(pk, &multisig.LegacyAminoPubKey{})

	var priv *cryptotypes.PrivKey
	registry.RegisterInterface("cosmos.crypto.PrivKey", priv)
	registry.RegisterImplementations(priv, &secp256k1.PrivKey{})
	registry.RegisterImplementations(priv, &ed25519.PrivKey{}) //nolint
	secp256r1.RegisterInterfaces(registry)
}
