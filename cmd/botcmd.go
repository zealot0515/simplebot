package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "casinobot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CasinoBot, args: ", args)
	},
}
