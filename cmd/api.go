package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Factom-Asset-Tokens/factom"
	"github.com/pegnet/pegnet/cmd"
	"github.com/pegnet/pegnetd/srv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(balances)
}

var balance = &cobra.Command{
	Use:              "balance",
	Short:            "",
	Example:          "pegnetd balance PEG FA2CEc2JSkhuckEXy42K111MvM9bycUDkbrrHjd9bNkBfvPBSGKd",
	PersistentPreRun: always,
	PreRun:           ReadConfig,
	Args: cmd.CombineCobraArgs(
		cmd.CustomArgOrderValidationBuilder(false, cmd.ArgValidatorAssetAndAll, cmd.ArgValidatorFCTAddress),
		cobra.MinimumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

var balances = &cobra.Command{
	Use:              "balances",
	Short:            "Fetch all balances for a given factoid address",
	Example:          "pegnetd balances FA2CEc2JSkhuckEXy42K111MvM9bycUDkbrrHjd9bNkBfvPBSGKd",
	PersistentPreRun: always,
	PreRun:           ReadConfig,
	Args: cmd.CombineCobraArgs(
		cmd.CustomArgOrderValidationBuilder(false, cmd.ArgValidatorFCTAddress),
		cobra.MinimumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		cl := srv.NewClient()
		// TODO: Able to change loc
		addr, err := factom.NewFAAddress(args[0])
		if err != nil {
			// TODO: Better error
			fmt.Println("1", err)
			os.Exit(1)
		}

		var res srv.ResultGetPegnetBalances
		err = cl.Request("get-pegnet-balances", srv.ParamsGetPegnetBalances{&addr}, &res)
		if err != nil {
			// TODO: Better error
			fmt.Println("2", err)
			os.Exit(1)
		}

		data, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	},
}