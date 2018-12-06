package states

import (
	"fmt"
	"simplebot/bot"
	"simplebot/utils/log"
	"tcpservertest/client/tcpclient"
)

// SocketTest tcpserver連線測試的state, 連線上後會持續送出cmd給server
type SocketTest struct {
	bot.BotState
	client   *tcpclient.TCPClient
	debugLog bool
}

func init() {
	bot.RegistState(&SocketTest{})

}

func (s *SocketTest) GetStateName() string {
	return "sockettest"
}

func (s *SocketTest) InitState(bot *bot.SimpleBotBase) {
	//會影響到下面那些LOG是否被印出
	s.debugLog = false
}

func (s *SocketTest) OnStateEnter(bot *bot.SimpleBotBase) {
	log.PrintDebugLog(s.debugLog, "sockettest.OnStateEnter ")
	var serverAddr = bot.IPPort
	s.client = tcpclient.Connect(serverAddr, s.onReceiveMsg)
}

func (s *SocketTest) OnStateExit(bot *bot.SimpleBotBase) {
	log.PrintDebugLog(s.debugLog, "sockettest.OnStateExit ")
}

func (s *SocketTest) OnUpdate(bot *bot.SimpleBotBase) {
	log.PrintDebugLog(s.debugLog, "sockettest.OnStateUpdate ")
	if s.client != nil {
		s.client.Send("queryapi1,param1,param2\n")
	}
}

func (s *SocketTest) OnForceStop(bot *bot.SimpleBotBase) {
	log.PrintDebugLog(s.debugLog, "sockettest.OnForceStop")
}

func (s *SocketTest) onReceiveMsg(receiveStr string) {
	fmt.Println("Got Server Msg:", receiveStr)
}
