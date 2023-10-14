package client

import (
	"github.com/opzlabs/cosmos-sdk/x/distribution/client/cli"
	govclient "github.com/opzlabs/cosmos-sdk/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal)
)
