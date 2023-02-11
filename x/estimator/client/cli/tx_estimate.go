package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"testhellochain/x/estimator/types"
)

var _ = strconv.Itoa(0)

func CmdEstimate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate [start] [end]",
		Short: "Broadcast message estimate",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argStart, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argEnd, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEstimate(
				clientCtx.GetFromAddress().String(),
				argStart,
				argEnd,
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
