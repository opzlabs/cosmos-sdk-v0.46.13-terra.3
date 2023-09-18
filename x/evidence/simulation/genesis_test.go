package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	sdkmath "cosmossdk.io/math"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec"
	codectypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec/types"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types/module"
	simtypes "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types/simulation"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/evidence/simulation"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/x/evidence/types"
)

// TestRandomizedGenState tests the normal scenario of applying RandomizedGenState.
// Abonormal scenarios are not tested here.
func TestRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: sdkmath.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)

	var evidenceGenesis types.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &evidenceGenesis)

	require.Len(t, evidenceGenesis.Evidence, 0)
}
