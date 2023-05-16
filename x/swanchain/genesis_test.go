package swanchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "swanchain/testutil/keeper"
	"swanchain/testutil/nullify"
	"swanchain/x/swanchain"
	"swanchain/x/swanchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SwanchainKeeper(t)
	swanchain.InitGenesis(ctx, *k, genesisState)
	got := swanchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
