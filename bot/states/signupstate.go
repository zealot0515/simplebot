package states

import (
	"casino_bot/bot"
	"casino_bot/util/signup"
	"casino_server/util/log"
)

// SignupBotState 註冊帳號的State, 如果帳號沒註冊過, 跑過這個State會進行註冊
type SignupBotState struct {
	bot.BotState
	debugLog bool
}

func init() {
	bot.RegistState(&SignupBotState{})

}

func (s *SignupBotState) GetStateName() string {
	return "signup"
}

func (s *SignupBotState) InitState(bot *bot.CasinoBot) {
	//會影響到下面那些LOG是否被印出
	s.debugLog = false
}

func (s *SignupBotState) OnStateEnter(bot *bot.CasinoBot) {
	log.PrintDebugLog(s.debugLog, "signup.OnStateEnter, ", bot.Account)
}

func (s *SignupBotState) OnStateExit(bot *bot.CasinoBot) {
	log.PrintDebugLog(s.debugLog, "signup.OnStateExit, ", bot.Account)
}

func (s *SignupBotState) OnUpdate(bot *bot.CasinoBot) {
	log.PrintDebugLog(s.debugLog, "signup.OnStateUpdate, ", bot.Account)
	signuputil.SignUp(bot.IPPort, bot.Account, bot.Password)
	bot.BotState.StateDone(s.GetStateName())
}

func (s *SignupBotState) OnForceStop(bot *bot.CasinoBot) {
	log.PrintDebugLog(s.debugLog, "signup.OnForceStop, ", bot.Account)
}
