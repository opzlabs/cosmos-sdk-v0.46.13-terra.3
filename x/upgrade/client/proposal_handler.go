package client

import (
	govclient "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/gov/client"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/upgrade/client/cli"
)

var (
	LegacyProposalHandler       = govclient.NewProposalHandler(cli.NewCmdSubmitLegacyUpgradeProposal)
	LegacyCancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitLegacyCancelUpgradeProposal)
)
