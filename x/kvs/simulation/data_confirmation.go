package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"kvs/x/kvs/keeper"
	"kvs/x/kvs/types"
)

func SimulateMsgDataConfirmation(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgDataConfirmation{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the DataConfirmation simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "DataConfirmation simulation not implemented"), nil, nil
	}
}
