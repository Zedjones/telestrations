package telestrationsLib

import (
	"errors"
	"time"
)

type playerState int

const (
	StateWorking playerState = 0
	StateDone    playerState = 1
	StateInLobby playerState = 2
)

type gameState int

const (
	StateProgress gameState = 0 // game in progress
	StateWaiting  gameState = 1 // hasnt started
	StateFinished gameState = 2 // all rounds over
)

type roundState int

const (
	RoundWaiting  roundState = 0
	RoundProgress roundState = 1
)

type Player struct {
	id    int
	name  string
	state playerState
}

type GameState struct {
	Round       int
	TimeLeft    int
	AllPlayers  map[int]Player
	state       gameState
	rState      roundState
	NumRounds   int
	RoundLength int
}

// --- GameState functions ---

func CreateGameState() *GameState {
	gm := new(GameState)
	gm.AllPlayers = make(map[int]Player)
	gm.state = StateWaiting
	return gm
}

func (gm GameState) GMAddPlayer(id int, name string) {
	gm.AllPlayers[id] = Player{id: id, name: name, state: StateInLobby}
}

func (gm GameState) GMSetState(state gameState) {
	gm.state = state
}

func (gm GameState) GMSetRoundLength(roundTime int) {
	gm.RoundLength = roundTime
}

func (gm GameState) GMGetPlayer(id int) *Player {
	play := gm.AllPlayers[id]
	return &play
}

func (gm GameState) GMGetPlayersByName(pname string) [](*Player) {
	players := make([](*Player), 0)
	for _, v := range gm.AllPlayers {
		if v.name == pname {
			players = append(players, &v)
		}
	}
	return players

}

func (gm GameState) GMGetPlayersAsArray() []Player {
	size := len(gm.AllPlayers)
	players := make([]Player, size)
	for i := 0; i < size; i++ {
		players[i] = gm.AllPlayers[i]
	}
	return players
}

func (gm GameState) GMStartGame() {
	gm.state = StateProgress
	//TODO
}

// Starts a new Round of the game
// Returns an error if all rounds completed
// returns nil otherwise
// spawns timer routine
// param roundTime the length of the round
func (gm GameState) GMRoundStart() error {
	gm.NumRounds = len(gm.AllPlayers)
	if gm.Round > gm.NumRounds {
		gm.rState = RoundWaiting
		gm.state = StateFinished
		return errors.New("All Rounds Over")
	}
	go gm.gmStartTimer()
	return nil

}

// Round Timer
func (gm GameState) gmStartTimer() {
	gm.rState = RoundProgress
	if gm.RoundLength == -1 { // for infinite time rounds
		for !gm.AllPlayersDone() {
			time.Sleep(time.Second)
		}
	} else { // for timed rounds
		gm.TimeLeft = gm.RoundLength
		for gm.TimeLeft > 0 {
			time.Sleep(time.Second)
			gm.TimeLeft--
		}
	}
	gm.Round++
	gm.rState = RoundWaiting
}

func (gm GameState) AllPlayersDone() bool {
	var b bool = true
	for _, v := range gm.AllPlayers {
		if v.state == StateWorking {
			b = false
		}
	}
	return b
}

// --- Player Functions ---

func (p Player) PLSetState(st playerState) {
	p.state = st
}

func (p Player) PLGetState() playerState {
	return p.state
}
