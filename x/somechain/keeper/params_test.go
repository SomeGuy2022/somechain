package keeper_test

import (
	"testing"

	testkeeper "github.com/SomeGuy2022/somechain/testutil/keeper"
	"github.com/SomeGuy2022/somechain/x/somechain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SomechainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
