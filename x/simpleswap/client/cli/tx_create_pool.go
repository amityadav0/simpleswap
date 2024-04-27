package cli

import (
	"fmt"
	"strconv"
	"strings"

	"simpleswap/x/simpleswap/types"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreatePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pool [tokens] [decimals] [swap-fee] ",
		Short: "Broadcast message create-pool",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			tokens, err := sdk.ParseCoinsNormalized(args[0])
			if err != nil {
				return err
			}

			decimals, err := parseDecimals(args[1])
			if err != nil {
				return err
			}

			swapFee, err := strconv.Atoi(args[2])
			if err != nil {
				return err
			}

			// swapFee is ranged from 0 to 10000.
			if swapFee > 10000 {
				return fmt.Errorf("swap fee must be less than 10000")
			}

			liquidity := []types.PoolAsset{}
			for i := 0; i < len(tokens); i++ {
				if len(tokens) != len(decimals) {
					return fmt.Errorf("tokens and decimals must have the same length")
				}
				liquidity = append(liquidity, types.PoolAsset{
					Token:   tokens[i],
					Decimal: sdk.NewInt(int64(decimals[i])),
				})
			}

			amp := math.NewInt(10)
			msg := types.NewMsgCreatePool(
				clientCtx.GetFromAddress().String(),
				types.PoolParams{
					SwapFee: sdk.NewDec(int64(swapFee)),
					Amp:     &amp,
				},
				liquidity,
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

func parseDecimals(decimalsStr string) ([]uint32, error) {
	decimalList := strings.Split(decimalsStr, ",")
	decimals := make([]uint32, 0, len(decimalList))

	for _, decimalStr := range decimalList {
		decimal, err := strconv.Atoi(decimalStr)
		if err != nil {
			return nil, fmt.Errorf("invalid decimal %s", decimalStr)
		}
		decimals = append(decimals, uint32(decimal))
	}

	if len(decimals) != 2 {
		return nil, fmt.Errorf("invalid decimals length %v", decimals)
	}

	return decimals, nil
}
