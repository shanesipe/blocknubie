package v2

import (
	"github.com/CosmosContracts/juno/v12/x/mint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "mint"
)

// Migrate migrates the x/mint module state from the consensus version 1 to
// version 2. Specifically, it take calculate target supply for the current phase
func Migrate(
	ctx sdk.Context,
	store sdk.KVStore,
	cdc codec.BinaryCodec,
) error {

	// Get minter
	var minter types.Minter
	b := store.Get(types.MinterKey)
	if b == nil {
		panic("stored minter should not have been nil")
	}

	cdc.MustUnmarshal(b, &minter)

	// Calculate target supply
	minter.TargetSupply = minter.AnnualProvisions.Add(minter.AnnualProvisions.Quo(minter.Inflation)).TruncateInt()

	// Save new minter
	bz := cdc.MustMarshal(&minter)
	store.Set(types.MinterKey, bz)

	return nil
}
