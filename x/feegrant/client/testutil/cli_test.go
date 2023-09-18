//go:build norace
// +build norace

package testutil

import (
	"testing"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/testutil/network"

	"github.com/stretchr/testify/suite"
)

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 3
	suite.Run(t, NewIntegrationTestSuite(cfg))
}
