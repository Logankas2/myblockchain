package myprojectpow_test

import (
	"testing"

	keepertest "github.com/username/myblockchain/testutil/keeper"
	"github.com/username/myblockchain/testutil/nullify"
	myprojectpow "github.com/username/myblockchain/x/myprojectpow/module"
	"github.com/username/myblockchain/x/myprojectpow/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MyprojectpowKeeper(t)
	myprojectpow.InitGenesis(ctx, k, genesisState)
	got := myprojectpow.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
