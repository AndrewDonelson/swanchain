package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/NlaakStudiosLLC/swan/testutil/keeper"
	"github.com/NlaakStudiosLLC/swanchain/x/swan/keeper"
	"github.com/NlaakStudiosLLC/swanchain/x/swan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SwanKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
