package cmd

import (
	"fmt"
	"simplebot/bot"

	"github.com/spf13/cobra"
)

var startbotCmd = &cobra.Command{
	Use:   "startbot",
	Short: "startbot command",
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 2:
			bot.StartBot(args[0], args[1])
		default:
			fmt.Printf("error args, use like this:startbot arg1,arg2,arg3 StateString")
		}
	},
}

func init() {
	RootCmd.AddCommand(startbotCmd)
}
