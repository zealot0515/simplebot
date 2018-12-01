package bot

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

var botPool []*SimpleBotBase

func init() {
	botPool = []*SimpleBotBase{}
}

// StartBot 啟動單一BOT
func StartBot(ipport string, acc string, pwd string, stateString string) {
	fmt.Printf("Start Bot, Acc:%s, stateString:%s...\n", acc, stateString)
	var bot = NewBot(ipport, acc, pwd, stateString)
	botPool = append(botPool, bot)
	bot.BotState.StartBot()
	fmt.Println("Start Bot Done")
	WaitMethod()
}

// StartMultiBot 啟動多個BOT
func StartMultiBot(ipport string, accPrefix string, pwd string, count int, stateString string) {
	fmt.Printf("Start Multi Bot, AccPrefix:%s, stateString:%s...\n", accPrefix, stateString)
	var bot *SimpleBotBase
	for i := 0; i < count; i++ {
		bot = NewBot(ipport, fmt.Sprintf("%s%v", accPrefix, i), pwd, stateString)
		botPool = append(botPool, bot)
		bot.BotState.StartBot()
		fmt.Printf("Start Bot %s Done\n", bot.Account)
	}
	WaitMethod()
}

func WaitMethod() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("Wait 5 Sec To Exit!")
	time.Sleep(time.Second * 5)
	fmt.Println("Exit!")
	os.Exit(1)
}
