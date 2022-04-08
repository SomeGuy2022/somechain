package keeper

import (
	"github.com/SomeGuy2022/somechain/x/somechain/types"
)

var _ types.QueryServer = Keeper{}
