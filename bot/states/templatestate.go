package states

import (
	"simplebot/bot"
)

// TemplateState 範例State
type TemplateState struct {
	bot.BotState
}

func init() {
	bot.RegistState(&TemplateState{})
}

func (s *TemplateState) GetStateName() string {
	return "template"
}

func (s *TemplateState) InitState(botObj interface{}) {
}

func (s *TemplateState) OnStateEnter(botObj interface{}) {
}

func (s *TemplateState) OnStateExit(botObj interface{}) {
}

func (s *TemplateState) OnUpdate(botObj interface{}) {
	botObj.(bot.SimpleBotBase).BotState.StateDone(s.GetStateName())
}

func (s *TemplateState) OnForceStop(botObj interface{}) {
}
