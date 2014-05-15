package request_handler

import (
	"api"
	"engine/game_engine"
	//"github.com/Tactique/golib/logger"
)

type Game struct {
	world *game_engine.World
}

func NewGame(request api.NewRequest) (*Game, error) {
	world, err := game_engine.NewWorld(request.Uids, request.Debug)
	return &Game{world: world}, err
}

func (game *Game) ViewWorld(playerId int, request api.ViewWorldRequest) (*api.ViewWorldResponse, error) {
	return game.world.ViewWorld(playerId)
}

func (game *Game) ViewTerrain(playerId int, request api.ViewTerrainRequest) (*api.ViewTerrainResponse, error) {
	return game.world.ViewTerrain(playerId)
}

func (game *Game) ViewPlayers(playerId int, request api.ViewPlayersRequest) (*api.ViewPlayersResponse, error) {
	return game.world.ViewPlayers(playerId)
}

func (game *Game) ViewUnits(playerId int, request api.ViewUnitsRequest) (*api.ViewUnitsResponse, error) {
	return game.world.ViewUnits(playerId)
}

func (game *Game) MoveUnit(playerId int, request api.MoveRequest) (*api.MoveResponse, error) {
	return game.world.MoveUnit(playerId, request.Move)
}

func (game *Game) Attack(playerId int, request api.AttackRequest) (*api.AttackResponse, error) {
	return game.world.Attack(playerId, request.Attacker, request.AttackIndex, request.Target)
}

func (game *Game) EndTurn(playerId int, request api.EndTurnRequest) (*api.EndTurnResponse, error) {
	return game.world.EndTurn(playerId)
}
