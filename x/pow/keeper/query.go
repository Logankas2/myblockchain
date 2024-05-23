package keeper

import (
	"github.com/username/myblockchain/x/pow/types"
)

var _ types.QueryServer = Keeper{}
