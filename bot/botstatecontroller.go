package bot

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"simplebot/utils/log"
	"strings"
	"time"

	"github.com/looplab/fsm"
)

type BotState interface {
	GetStateName() string
	InitState(botObj interface{})
	OnStateEnter(botObj interface{})
	OnStateExit(botObj interface{})
	OnUpdate(botObj interface{})
	OnForceStop(botObj interface{})
}

type BotStateController struct {
	bot          interface{}
	stateMap     map[string]BotState
	currentState string
	botFsm       *fsm.FSM
	ticker       *time.Ticker
	Signal       chan os.Signal
	tickerDur    time.Duration
	firstState   string
	debugLog     bool
}

var botStatesMap map[string]BotState
var defaultstateUpdateDuration = time.Duration(100 * time.Millisecond)
var initState = "wait"

func RegistState(state BotState) {
	var name = state.GetStateName()
	if botStatesMap == nil {
		botStatesMap = map[string]BotState{}
	}
	if _, ok := botStatesMap[name]; ok {
		fmt.Printf("error state[%s] exist, can't regist!\n", name)
	} else {
		botStatesMap[name] = state
		fmt.Printf("State[%s] Regist!\n", name)
	}
}

func NewBotStateController(bot interface{}, stateString string) (controller *BotStateController) {
	controller = &BotStateController{}
	controller.Init(bot, stateString)
	return controller
}

func (b *BotStateController) Init(bot interface{}, stateString string) {
	b.bot = bot
	b.tickerDur = bot.(BotBaseGetter).GetHeartBeatTime()
	if b.tickerDur <= 0 {
		b.tickerDur = defaultstateUpdateDuration
	}
	b.debugLog = false
	//Load All State And Init Fsm
	b.initFSM(stateString)
	go b.startTicker()
}

func (b *BotStateController) StartBot() {
	log.PrintDebugLog(b.debugLog, "Let Bot To initState: ", initState)
	err := b.botFsm.Event(initState)
	if err != nil {
		fmt.Println("err:", err)
	}
}

func (b *BotStateController) StateDone(stateString string) {
	log.PrintDebugLog(b.debugLog, "BotStateController.StateDone:", stateString, b.botFsm)
	var err error
	err = b.botFsm.Event(stateString)
	if err != nil {
		fmt.Println("change state error:", err)
	}
}

func (b *BotStateController) onStateChangeTo(e *fsm.Event) {
	var stateName = e.FSM.Current()
	log.PrintDebugLog(b.debugLog, "BotStateController.onStateChangeTo:", stateName)
	if state, ok := b.stateMap[stateName]; ok {
		state.OnStateEnter(b.bot)
	}
	b.currentState = stateName
}

func (b *BotStateController) onStateLeaveFrom(e *fsm.Event) {
	var stateName = e.FSM.Current()
	log.PrintDebugLog(b.debugLog, "BotStateController.onStateLeaveFrom:", stateName)
	if state, ok := b.stateMap[stateName]; ok {
		state.OnStateExit(b.bot)
	}
}

func (b *BotStateController) onStateUpdate() {
	log.PrintDebugLog(b.debugLog, "BotStateController.onStateUpdate:", b.currentState)
	if state, ok := b.stateMap[b.currentState]; ok {
		state.OnUpdate(b.bot)
	}
}

func (b *BotStateController) onForceStop() {
	log.PrintDebugLog(b.debugLog, "BotStateController.onForceStop:", b.currentState)
	if state, ok := b.stateMap[b.currentState]; ok {
		state.OnForceStop(b.bot)
	}
}

func (b *BotStateController) startTicker() {
	b.ticker = time.NewTicker(b.tickerDur)
	b.Signal = make(chan os.Signal, 1)
	signal.Notify(b.Signal, os.Interrupt)
	for {
		select {
		case <-b.ticker.C:
			b.onStateUpdate()
		case <-b.Signal:
			b.ticker.Stop()
			b.onForceStop()
			return
		}
	}
}

// initFSM init States And Create FSM Event
func (b *BotStateController) initFSM(stateString string) {
	b.stateMap = map[string]BotState{}
	var stateNames = strings.Split(stateString, ",")
	if len(stateNames) > 0 {
		var fsmEvents = fsm.Events{}
		b.firstState = stateNames[0]
		fmt.Println(botStatesMap)
		fsmEvents = append(fsmEvents, fsm.EventDesc{Name: initState, Src: []string{initState}, Dst: b.firstState})
		for idx, stateName := range stateNames {
			fmt.Println(stateName)
			if state, ok := botStatesMap[stateName]; ok {
				b.stateMap[stateName] = reflect.New(reflect.ValueOf(state).Elem().Type()).Interface().(BotState)
				b.stateMap[stateName].InitState(b.bot)
				if idx+1 >= len(stateNames) {
					fsmEvents = append(fsmEvents, fsm.EventDesc{Name: stateName, Src: []string{stateName}, Dst: b.firstState})
				} else {
					fsmEvents = append(fsmEvents, fsm.EventDesc{Name: stateName, Src: []string{stateName}, Dst: stateNames[idx+1]})
				}
			}
		}
		b.botFsm = fsm.NewFSM(
			initState,
			fsmEvents,
			fsm.Callbacks{
				"enter_state": func(e *fsm.Event) { b.onStateChangeTo(e) },
				"leave_state": func(e *fsm.Event) { b.onStateLeaveFrom(e) },
			},
		)
	}
}
