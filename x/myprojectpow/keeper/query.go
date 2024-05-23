package keeper

import (
	"github.com/username/myblockchain/x/myprojectpow/types"
)

var _ types.QueryServer = Keeper{}
