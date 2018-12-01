package cmd

import (
	"casino_bot/bot"
	"fmt"

	"github.com/spf13/cobra"
)

var startbotCmd = &cobra.Command{
	Use:   "startbot",
	Short: "startbot command",
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 4:
			bot.StartBot(args[0], args[1], args[2], args[3])
		default:
			fmt.Printf("error args, use like this:startbot ip:port Acc Pwd StateString")
		}
	},
}

func init() {
	RootCmd.AddCommand(startbotCmd)
}
