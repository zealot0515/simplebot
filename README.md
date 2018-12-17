## 這是一個用GO建立的簡易BOT框架
##### 使用命令列套件(github.com/spf13/cobra)與狀態機套件(github.com/looplab/fsm)
使用方式:
- 建立自己的bot物件
- 建立自己的bot states

### 命令列使用方式
``` 
單一bot
simplebot cmd paramstr statestr
paramstr->參數字串,多個以逗號分隔
statestr->狀態機描述字串, 多個以逗號分隔, 狀態會循序切換, 再返回第一個

ex:
simplebot startbot 111,222,333,444,555 hello
```
```
複數bot
simplebot cmd paramstr statestr count
paramstr->參數字串,多個以逗號分隔
statestr->狀態機描述字串, 多個以逗號分隔, 狀態會循序切換, 再返回第一個
count->bot開啟數量

ex:
simplebot startmultibot 111,222,333,444,555 hello 10
```
### 建立bot物件 (可參考example目錄底下的hellobot)
``` go
package hellobot

import (
	"simplebot/bot"
	"strings"
	"time"
)

type HelloBot struct {
	bot.SimpleBotBase  <=必須繼承的struct
	Messages []string <=bot自己的資料, 可任意新增
	bot.BotCreator <=必須建立一個產生bot的接口
}

//初始化並回傳一個bot物件
func (b *HelloBot) New(paramString string, botIdx int) interface{} {
	var rtnBot = &HelloBot{}
	rtnBot.Messages = strings.Split(paramString, ",")
	rtnBot.BotHeartBeatTime = time.Second  <=bot心跳週期, 不給值得話會是100ms
	return rtnBot
}

func init() {
	bot.RegistCreator(&HelloBot{})
}
```
### 建立bot state (可參考example目錄底下的hellobot)
``` go
package hellostates

import (
	"fmt"
	"simplebot/bot"
	"simplebot/example/hellobot"
)

// HelloState Print Param strings
type HelloState struct { //這邊狀態機有需求也可以加入自己的資料結構
	bot.BotState
}

func init() {
	bot.RegistState(&HelloState{})
}

func (s *HelloState) GetStateName() string { //狀態名稱, 對應到指令列要輸入的狀態字串
	return "hello"
}

func (s *HelloState) InitState(botObj interface{}) { //狀態機初始化
}

func (s *HelloState) OnStateEnter(botObj interface{}) { //進入狀態
}

func (s *HelloState) OnStateExit(botObj interface{}) { //離開狀態
}

func (s *HelloState) OnUpdate(botObj interface{}) {  //狀態機心跳會定時觸發的
	var helloBot = botObj.(*hellobot.HelloBot)
	if len(helloBot.Messages) > 0 {
		fmt.Println(helloBot.Messages[0])
		helloBot.Messages = helloBot.Messages[1:]
	} else {
		helloBot.BotState.StateDone(s.GetStateName())  <==切換到下一個state
	}
}

func (s *HelloState) OnForceStop(botObj interface{}) { //命令列收到Ctrl+c
}
```
### main.go新增預讀
``` go
package main

import (
	"fmt"
	"os"
	"simplebot/cmd"
	_ "simplebot/example/hellobot"          <= 預載bot
	_ "simplebot/example/hellobot/states"   <= 預載bot state
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

### License
```
SimpleBot is licensed under Apache License 2.0
http://www.apache.org/licenses/LICENSE-2.0
```
