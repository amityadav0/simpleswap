package simpleswap

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"simpleswap/testutil/sample"
	simpleswapsimulation "simpleswap/x/simpleswap/simulation"
	"simpleswap/x/simpleswap/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = simpleswapsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreatePool = "op_weight_msg_create_pool"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePool int = 100

	opWeightMsgAddLiquidity = "op_weight_msg_add_liquidity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddLiquidity int = 100

	opWeightMsgWithdraw = "op_weight_msg_withdraw"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdraw int = 100

	opWeightMsgSwap = "op_weight_msg_swap"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSwap int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	simpleswapGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&simpleswapGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreatePool, &weightMsgCreatePool, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePool = defaultWeightMsgCreatePool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePool,
		simpleswapsimulation.SimulateMsgCreatePool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddLiquidity int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddLiquidity, &weightMsgAddLiquidity, nil,
		func(_ *rand.Rand) {
			weightMsgAddLiquidity = defaultWeightMsgAddLiquidity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddLiquidity,
		simpleswapsimulation.SimulateMsgAddLiquidity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdraw int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdraw, &weightMsgWithdraw, nil,
		func(_ *rand.Rand) {
			weightMsgWithdraw = defaultWeightMsgWithdraw
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdraw,
		simpleswapsimulation.SimulateMsgWithdraw(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSwap int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSwap, &weightMsgSwap, nil,
		func(_ *rand.Rand) {
			weightMsgSwap = defaultWeightMsgSwap
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSwap,
		simpleswapsimulation.SimulateMsgSwap(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatePool,
			defaultWeightMsgCreatePool,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				simpleswapsimulation.SimulateMsgCreatePool(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAddLiquidity,
			defaultWeightMsgAddLiquidity,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				simpleswapsimulation.SimulateMsgAddLiquidity(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgWithdraw,
			defaultWeightMsgWithdraw,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				simpleswapsimulation.SimulateMsgWithdraw(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSwap,
			defaultWeightMsgSwap,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				simpleswapsimulation.SimulateMsgSwap(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
