package client

import (
	govclient "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/gov/client"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/params/client/cli"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd)
