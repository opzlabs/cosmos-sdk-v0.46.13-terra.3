package rosetta

import (
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec"
	codectypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec/types"
	cryptocodec "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/crypto/codec"
	authcodec "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/auth/types"
	bankcodec "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/bank/types"
)

// MakeCodec generates the codec required to interact
// with the cosmos APIs used by the rosetta gateway
func MakeCodec() (*codec.ProtoCodec, codectypes.InterfaceRegistry) {
	ir := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(ir)

	authcodec.RegisterInterfaces(ir)
	bankcodec.RegisterInterfaces(ir)
	cryptocodec.RegisterInterfaces(ir)

	return cdc, ir
}
