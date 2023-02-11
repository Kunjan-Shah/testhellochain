package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"testhellochain/x/estimator/types"
)

func CmdListApiCountMap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-api-count-map",
		Short: "list all apiCountMap",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllApiCountMapRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ApiCountMapAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowApiCountMap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-api-count-map [creator]",
		Short: "shows a apiCountMap",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argCreator := args[0]

			params := &types.QueryGetApiCountMapRequest{
				Creator: argCreator,
			}

			res, err := queryClient.ApiCountMap(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
