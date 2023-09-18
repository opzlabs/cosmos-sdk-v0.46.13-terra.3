package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	sdkmath "cosmossdk.io/math"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/simapp"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types/module"
	simtypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types/simulation"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/nft"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/nft/simulation"
)

func TestRandomizedGenState(t *testing.T) {
	app := simapp.Setup(t, false)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          app.AppCodec(),
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: sdkmath.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)
	var nftGenesis nft.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[nft.ModuleName], &nftGenesis)

	require.Len(t, nftGenesis.Classes, len(simState.Accounts)-1)
	require.Len(t, nftGenesis.Entries, len(simState.Accounts)-1)
}
