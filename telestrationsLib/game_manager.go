package telestrationsLib

import (
	"errors"
	"sort"
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
	ID    int
	Name  string
	State playerState
}

type GameState struct {
	Round       int
	TimeLeft    int
	AllPlayers  map[int]Player
	State       gameState
	RState      roundState
	NumRounds   int
	RoundLength int
}

// --- GameState functions ---

func CreateGameState() *GameState {
	gm := new(GameState)
	gm.AllPlayers = make(map[int]Player)
	gm.State = StateWaiting
	return gm
}

func (gm GameState) GMAddPlayer(id int, name string) {
	gm.AllPlayers[id] = Player{ID: id, Name: name, State: StateInLobby}
}

func (gm GameState) GMSetState(state gameState) {
	gm.State = state
}

func (gm GameState) GMSetRoundLength(roundTime int) {
	gm.RoundLength = roundTime
}

func (gm GameState) GMGetPlayer(id int) *Player {
	for _, v := range gm.AllPlayers {
		if v.ID == id {
			return &v
		}
	}
	// play := gm.AllPlayers[id]
	return &Player{999, "aaaa", StateDone}
}

func (gm GameState) GMPlayerExists(id int) bool {
	var b bool = false
	for _, v := range gm.AllPlayers {
		if v.ID == id {
			b = true
			break
		}
	}
	return b
}

func (gm GameState) GMGetPlayersByName(pname string) [](*Player) {
	players := make([](*Player), 0)
	for _, v := range gm.AllPlayers {
		if v.Name == pname {
			players = append(players, &v)
		}
	}
	return players

}

func (gm GameState) GMGetPlayersAsArray() []Player {
	players := []Player{}
	for _, player := range gm.AllPlayers {
		players = append(players, player)
	}
	sort.Slice(players, func(i, j int) bool {
		return players[i].ID < players[j].ID
	})
	return players
}

// @pre all players have joined / been added
// @pre round length has been set
func (gm GameState) GMStartGame() {
	gm.State = StateProgress
	go gm.GMRoundStart()
}

// Starts a new Round of the game
// Returns an error if all rounds completed
// returns nil otherwise
// spawns timer routine
// param roundTime the length of the round
func (gm GameState) GMRoundStart() error {
	gm.NumRounds = len(gm.AllPlayers)
	for {
		if gm.Round > gm.NumRounds {
			gm.RState = RoundWaiting
			gm.State = StateFinished
			return errors.New("All Rounds Over")
		}
		gm.gmStartTimer()
	}
}

// Round Timer
func (gm GameState) gmStartTimer() {
	gm.RState = RoundProgress
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
	gm.RState = RoundWaiting
}

func (gm GameState) AllPlayersDone() bool {
	var b bool = true
	for _, v := range gm.AllPlayers {
		if v.State == StateWorking {
			b = false
		}
	}
	return b
}

// --- Player Functions ---

func (p Player) PLSetState(st playerState) {
	p.State = st
}

func (p Player) PLGetState() playerState {
	return p.State
}
