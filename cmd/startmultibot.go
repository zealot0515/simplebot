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
		case 3:
			count, err = strconv.Atoi(args[2])
			if err != nil {
				fmt.Println("get count error:", err)
				return
			}

			bot.StartMultiBot(args[0], args[1], count)
		default:
			fmt.Printf("error args, use like this:startmultibot arg1,arg2,arg3 StateString botcount")
		}
	},
}

func init() {
	RootCmd.AddCommand(startmultibotCmd)
}
