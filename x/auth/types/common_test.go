package types_test

import (
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/simapp"
)

var (
	ecdc                  = simapp.MakeTestEncodingConfig()
	appCodec, legacyAmino = ecdc.Codec, ecdc.Amino
)
