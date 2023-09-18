package distribution

import (
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	sdkerrors "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types/errors"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/distribution/keeper"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/distribution/types"
	govtypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/gov/types/v1beta1"
)

func NewCommunityPoolSpendProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.CommunityPoolSpendProposal:
			return keeper.HandleCommunityPoolSpendProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized distr proposal content type: %T", c)
		}
	}
}
