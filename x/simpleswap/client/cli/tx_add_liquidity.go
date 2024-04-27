package cli

import (
	"strconv"

	"simpleswap/x/simpleswap/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddLiquidity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-liquidity [pool-id] [liquidity]",
		Short: "Broadcast message add-liquidity",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPoolId := args[0]

			argLiquidity, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddLiquidity(
				clientCtx.GetFromAddress().String(),
				argPoolId,
				argLiquidity,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
