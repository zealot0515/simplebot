package bot

var botCreator BotCreator

// NewBot 產一個新的BOT實體
func NewBot(paramString string, stateString string, botIdx int) (bot interface{}) {
	if botCreator != nil {
		bot = botCreator.New(paramString, botIdx)
		var botBase = bot.(BotBaseGetter).GetBotBase()
		botBase.BotState = NewBotStateController(bot, stateString)
	}
	return bot
}

func RegistCreator(creator BotCreator) {
	botCreator = creator
}
