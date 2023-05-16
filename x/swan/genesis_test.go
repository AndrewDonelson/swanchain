package swan_test

import (
	"testing"

	keepertest "github.com/NlaakStudiosLLC/swan/testutil/keeper"
	"github.com/NlaakStudiosLLC/swan/testutil/nullify"
	"github.com/NlaakStudiosLLC/swan/x/swan"
	"github.com/NlaakStudiosLLC/swanchain/x/swan/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SwanKeeper(t)
	swan.InitGenesis(ctx, *k, genesisState)
	got := swan.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
