package cli_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/network"

	"github.com/stretchr/testify/suite"

	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
)

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 1

	suite.Run(t, banktestutil.NewIntegrationTestSuite(cfg))
}
