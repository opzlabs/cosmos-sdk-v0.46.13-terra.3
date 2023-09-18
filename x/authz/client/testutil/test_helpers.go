package testutil

import (
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/testutil"
	clitestutil "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/testutil/cli"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/testutil/network"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/authz/client/cli"
)

func CreateGrant(val *network.Validator, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization()
	clientCtx := val.ClientCtx
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
