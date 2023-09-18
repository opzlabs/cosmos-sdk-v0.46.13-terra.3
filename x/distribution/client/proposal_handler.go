package client

import (
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/distribution/client/cli"
	govclient "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal)
)
