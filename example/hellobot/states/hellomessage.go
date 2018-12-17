package hellostates

import (
	"fmt"
	"simplebot/bot"
	"simplebot/example/hellobot"
)

// HelloState Print Param strings
type HelloState struct {
	bot.BotState
}

func init() {
	bot.RegistState(&HelloState{})
}

func (s *HelloState) GetStateName() string {
	return "hello"
}

func (s *HelloState) InitState(botObj interface{}) {
}

func (s *HelloState) OnStateEnter(botObj interface{}) {
}

func (s *HelloState) OnStateExit(botObj interface{}) {
}

func (s *HelloState) OnUpdate(botObj interface{}) {
	var helloBot = botObj.(*hellobot.HelloBot)
	if len(helloBot.Messages) > 0 {
		fmt.Println(helloBot.Messages[0])
		helloBot.Messages = helloBot.Messages[1:]
	} else {
		helloBot.BotState.StateDone(s.GetStateName())
	}
}

func (s *HelloState) OnForceStop(botObj interface{}) {
}
