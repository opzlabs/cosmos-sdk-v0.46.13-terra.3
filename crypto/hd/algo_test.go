package hd_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/crypto/hd"
)

func TestDefaults(t *testing.T) {
	require.Equal(t, hd.PubKeyType("multi"), hd.MultiType)
	require.Equal(t, hd.PubKeyType("secp256k1"), hd.Secp256k1Type)
	require.Equal(t, hd.PubKeyType("ed25519"), hd.Ed25519Type)
	require.Equal(t, hd.PubKeyType("sr25519"), hd.Sr25519Type)
}
