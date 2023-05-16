package keeper

import (
	"swanchain/x/swanchain/types"
)

var _ types.QueryServer = Keeper{}
