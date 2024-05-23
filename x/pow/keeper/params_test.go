package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

    keepertest "github.com/username/myblockchain/testutil/keeper"
    "github.com/username/myblockchain/x/pow/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.PowKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
