package kvs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"kvs/x/kvs/keeper"
	"kvs/x/kvs/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the data
	for _, elem := range genState.DataList {
		k.SetData(ctx, elem)
	}
	// Set all the proposal
	for _, elem := range genState.ProposalList {
		k.SetProposal(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.DataList = k.GetAllData(ctx)
	genesis.ProposalList = k.GetAllProposal(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
