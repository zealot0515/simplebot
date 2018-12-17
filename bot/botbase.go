package bot

import "time"

type SimpleBotBase struct {
	BotHeartBeatTime time.Duration
	BotState         *BotStateController
}

type BotCreator interface {
	New(paramString string, botIdx int) interface{}
}

type BotBaseGetter interface {
	GetBotBase() *SimpleBotBase
	GetHeartBeatTime() time.Duration
}

func (b *SimpleBotBase) GetBotBase() *SimpleBotBase {
	return b
}

func (b *SimpleBotBase) GetHeartBeatTime() time.Duration {
	return b.BotHeartBeatTime
}
