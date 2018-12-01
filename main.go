package main

import (
	_ "casino_bot/bot"
	_ "casino_bot/bot/states"
	"casino_bot/cmd"
	"fmt"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
