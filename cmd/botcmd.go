package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "simplebot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SimpleBotBase, args: ", args)
	},
}
