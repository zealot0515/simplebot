package bot

type SimpleBotBase struct {
	IPPort   string
	Account  string
	Password string
	BotData  BotDatas
	BotState *BotStateController
}

type BotDatas struct {
	PlayMoney int64
}
