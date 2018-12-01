package cmd

import (
	"fmt"
	"simplebot/bot"
	"strconv"

	"github.com/spf13/cobra"
)

var startmultibotCmd = &cobra.Command{
	Use:   "startmultibot",
	Short: "startmultibot command",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var count int
		switch len(args) {
		case 5:
			count, err = strconv.Atoi(args[3])
			if err != nil {
				fmt.Println("get count error:", err)
				return
			}

			bot.StartMultiBot(args[0], args[1], args[2], count, args[4])
		default:
			fmt.Printf("error args, use like this:startmultibot ip:port Acc Pwd count StateString")
		}
	},
}

func init() {
	RootCmd.AddCommand(startmultibotCmd)
}
