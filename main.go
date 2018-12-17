package main

import (
	"fmt"
	"os"
	"simplebot/cmd"
	_ "simplebot/example/hellobot"
	_ "simplebot/example/hellobot/states"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
