package keeper_test

import (
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/simapp"
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
	authtypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/auth/types"
)

// returns context and app with params set on account keeper
func createTestApp(t *testing.T, isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(t, isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())

	return app, ctx
}
