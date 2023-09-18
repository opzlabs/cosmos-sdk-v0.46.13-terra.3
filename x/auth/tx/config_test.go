package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec"
	codectypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec/types"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/std"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/testutil/testdata"
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
