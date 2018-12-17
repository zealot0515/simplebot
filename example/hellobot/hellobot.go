package hellobot

import (
	"simplebot/bot"
	"strings"
	"time"
)

type HelloBot struct {
	bot.SimpleBotBase
	Messages []string
	bot.BotCreator
}

func (b *HelloBot) New(paramString string, botIdx int) interface{} {
	var rtnBot = &HelloBot{}
	rtnBot.Messages = strings.Split(paramString, ",")
	rtnBot.BotHeartBeatTime = time.Second
	return rtnBot
}

func init() {
	bot.RegistCreator(&HelloBot{})
}
