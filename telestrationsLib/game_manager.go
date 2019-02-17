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
	AllPlayers map[string]Player
}

func CreateGameState() *GameState {
	gameState := new(GameState)
	gameState.AllPlayers = make(map[string]Player)
	return gameState
}

func (gm GameState) GMAddPlayer(id int, name string) {
	gm.AllPlayers[name] = Player{id: len(gm.AllPlayers), name: name, state: StateInLobby}
}