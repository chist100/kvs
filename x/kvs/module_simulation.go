package kvs

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"kvs/testutil/sample"
	kvssimulation "kvs/x/kvs/simulation"
	"kvs/x/kvs/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = kvssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgDataProposal = "op_weight_msg_data_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDataProposal int = 100

	opWeightMsgDataConfirmation = "op_weight_msg_data_confirmation"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDataConfirmation int = 100

	opWeightMsgAddressRegistration = "op_weight_msg_address_registration"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddressRegistration int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	kvsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&kvsGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgDataProposal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDataProposal, &weightMsgDataProposal, nil,
		func(_ *rand.Rand) {
			weightMsgDataProposal = defaultWeightMsgDataProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDataProposal,
		kvssimulation.SimulateMsgDataProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDataConfirmation int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDataConfirmation, &weightMsgDataConfirmation, nil,
		func(_ *rand.Rand) {
			weightMsgDataConfirmation = defaultWeightMsgDataConfirmation
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDataConfirmation,
		kvssimulation.SimulateMsgDataConfirmation(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddressRegistration int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddressRegistration, &weightMsgAddressRegistration, nil,
		func(_ *rand.Rand) {
			weightMsgAddressRegistration = defaultWeightMsgAddressRegistration
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddressRegistration,
		kvssimulation.SimulateMsgAddressRegistration(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
