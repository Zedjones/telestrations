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
	StateProgress gameState = 0 // round in progress
	StateEmpty    gameState = 1 // has not started
	StateWaiting  gameState = 2 // in between rounds
	StateFinished gameState = 3 // all rounds over
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
	NumRounds   int
	RoundLength int
}

// --- GameState functions ---

func CreateGameState() *GameState {
	gm := new(GameState)
	gm.AllPlayers = make(map[int]Player)
	gm.state = StateEmpty
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

// sets GameState instance to finished
func (gm GameState) gmFinished() {
	gm.state = StateFinished
}

// Starts a new Round of the game
// Returns an error if all rounds completed
// returns nil otherwise
// spawns timer routine
// param roundTime the length of the round
func (gm GameState) GMRoundStart() error {
	gm.NumRounds = len(gm.AllPlayers)
	if gm.Round > gm.NumRounds {
		gm.gmFinished()
		return errors.New("All Rounds Over")
	}
	go gm.gmStartTimer()
	return nil

}

// Round Timer
func (gm GameState) gmStartTimer() {
	gm.state = StateProgress
	if gm.RoundLength == -1 {

	} else {
		gm.TimeLeft = gm.RoundLength
		for gm.TimeLeft > 0 {
			time.Sleep(time.Second)
			gm.TimeLeft--
		}
	}
	gm.Round++
	gm.state = StateWaiting
}

func (gm GameState) allPlayersDone() bool {
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
