package testutil

import (
	"github.com/opzlabs/cosmos-sdk/testutil"
	clitestutil "github.com/opzlabs/cosmos-sdk/testutil/cli"
	"github.com/opzlabs/cosmos-sdk/testutil/network"
	"github.com/opzlabs/cosmos-sdk/x/authz/client/cli"
)

func CreateGrant(val *network.Validator, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization()
	clientCtx := val.ClientCtx
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
