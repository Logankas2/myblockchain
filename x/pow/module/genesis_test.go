package pow_test

import (
	"testing"

	keepertest "github.com/username/myblockchain/testutil/keeper"
	"github.com/username/myblockchain/testutil/nullify"
	pow "github.com/username/myblockchain/x/pow/module"
	"github.com/username/myblockchain/x/pow/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PowKeeper(t)
	pow.InitGenesis(ctx, k, genesisState)
	got := pow.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
