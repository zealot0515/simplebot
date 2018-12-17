package bot

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

var botPool []interface{}

func init() {
	botPool = []interface{}{}
}

// StartBot 啟動單一BOT
func StartBot(paramString string, stateString string) {
	fmt.Printf("Start Bot, paramString:%s, stateString:%s...\n", paramString, stateString)
	var bot = NewBot(paramString, stateString, 0)
	botPool = append(botPool, bot)
	bot.(BotBaseGetter).GetBotBase().BotState.StartBot()
	fmt.Println("Start Bot Done")
	WaitMethod()
}

// StartMultiBot 啟動多個BOT
func StartMultiBot(paramString string, stateString string, count int) {
	var bot interface{}
	for i := 0; i < count; i++ {
		bot = NewBot(paramString, stateString, i)
		botPool = append(botPool, bot)
		bot.(BotBaseGetter).GetBotBase().BotState.StartBot()
		fmt.Printf("Start Bot %d Done\n", i)
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
