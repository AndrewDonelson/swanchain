package keeper_test

import (
	"testing"

	testkeeper "github.com/NlaakStudiosLLC/swan/testutil/keeper"
	"github.com/NlaakStudiosLLC/swanchain/x/swan/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SwanKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
