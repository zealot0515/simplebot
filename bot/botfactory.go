package bot

// NewBot 產一個新的BOT實體
func NewBot(ipport string, acc string, pwd string, stateString string) (bot *CasinoBot) {
	bot = &CasinoBot{
		IPPort:   ipport,
		Account:  acc,
		Password: pwd,
		BotData: BotDatas{
			PlayMoney: 0,
		},
		BotState: nil,
	}
	bot.BotState = NewBotStateController(bot, stateString)
	return bot
}
