package keeper

import (
	"crudechain/x/crude/types"
)

var _ types.QueryServer = Keeper{}
