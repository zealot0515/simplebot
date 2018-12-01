package main

import (
	"fmt"
	"os"
	_ "simplebot/bot"
	_ "simplebot/bot/states"
	"simplebot/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}