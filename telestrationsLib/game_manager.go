package telestrationsLib

type playerState int

const (
	StateWorking playerState = 0
	StateDone    playerState = 1
	StateInLobby playerState = 2
)

type gameState int

const (
	StateProgress gameState = 0
	StateEmpty    gameState = 1
	StateWaiting  gameState = 2
)

type Player struct {
	id    int
	name  string
	state playerState
}

type GameState struct {
	Round      int
	TimeLeft   int
	AllPlayers map[int]Player
	state      gameState
}

func CreateGameState() *GameState {
	gameState := new(GameState)
	gameState.AllPlayers = make(map[int]Player)
	return gameState
}

func (gm GameState) GMAddPlayer(id int, name string) {
	gm.AllPlayers[id] = Player{id: id, name: name, state: StateInLobby}
}

func (gm GameState) GMSetState(state gameState) {
	gm.state = state
}

func (gm GameState) GMGetPlayer(id int) {

}

// Round Timer
// GetPlayer
