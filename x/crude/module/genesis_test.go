package crude_test

import (
	"testing"

	keepertest "crudechain/testutil/keeper"
	"crudechain/testutil/nullify"
	"crudechain/x/crude/module"
	"crudechain/x/crude/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CrudeKeeper(t)
	crude.InitGenesis(ctx, k, genesisState)
	got := crude.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
