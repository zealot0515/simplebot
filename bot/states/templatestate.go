package states

import (
	"simplebot/bot"
	"simplebot/utils/log"
)

// TemplateState 註冊帳號的State, 如果帳號沒註冊過, 跑過這個State會進行註冊
type TemplateState struct {
	bot.BotState
	debugLog bool
}

func init() {
	bot.RegistState(&TemplateState{})

}

func (s *TemplateState) GetStateName() string {
	return "template"
}

func (s *TemplateState) InitState(bot *bot.SimpleBotBase) {
	//會影響到下面那些LOG是否被印出
	s.debugLog = false
}

func (s *TemplateState) OnStateEnter(bot *bot.SimpleBotBase) {
	log.PrintDebugLog(s.debugLog, "template.OnStateEnter")
}

func (s *TemplateState) OnStateExit(bot *bot.SimpleBotBase) {
	log.PrintDebugLog(s.debugLog, "template.OnStateExit")
}

func (s *TemplateState) OnUpdate(bot *bot.SimpleBotBase) {
	log.PrintDebugLog(s.debugLog, "template.OnStateUpdate")

	bot.BotState.StateDone(s.GetStateName())
}

func (s *TemplateState) OnForceStop(bot *bot.SimpleBotBase) {
	log.PrintDebugLog(s.debugLog, "template.OnForceStop")
}
