package bot

import (
	"strings"
)

// NewBot 產一個新的BOT實體
func NewBot(paramString string, stateString string, botIdx int) (bot *SimpleBotBase) {
	var params = strings.Split(paramString, ",")
	bot = &SimpleBotBase{
		IPPort:   params[0],
		BotData:  BotDatas{},
		BotState: nil,
	}
	bot.BotState = NewBotStateController(bot, stateString)
	return bot
}
