package request_handler

import (
	"api"
	"fmt"
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
	terrain, err := game.ViewTerrain(playerId, api.ViewTerrainRequest{})
	if err != nil {
		return nil, err
	}
	players, err := game.ViewPlayers(playerId, api.ViewPlayersRequest{})
	if err != nil {
		return nil, err
	}
	units, err := game.ViewUnits(playerId, api.ViewUnitsRequest{})
	if err != nil {
		return nil, err
	}
	return &api.ViewWorldResponse{
		TerrainResponse: terrain,
		UnitsResponse:   units,
		PlayersResponse: players}, nil
}

func (game *Game) ViewTerrain(playerId int, request api.ViewTerrainRequest) (*api.ViewTerrainResponse, error) {
	terrain := game.world.GetTerrain()
	terrainInts := make([][]int, len(terrain))
	for i, t := range terrain {
		thoriz := make([]int, len(t))
		for j, t_ := range t {
			thoriz[j] = int(t_)
		}
		terrainInts[i] = thoriz
	}
	return &api.ViewTerrainResponse{
		Terrain: terrainInts}, nil
}

func (game *Game) ViewPlayers(playerId int, request api.ViewPlayersRequest) (*api.ViewPlayersResponse, error) {
	players := make(map[string]*api.PlayerStruct, game.world.GetNumPlayers())
	for _, player := range game.world.GetPlayers() {
		players[fmt.Sprintf("%d", player.GetPlayerId())] = player.Serialize()
	}
	return &api.ViewPlayersResponse{
		Players:   players,
		TurnOwner: int(game.world.GetTurnOwner().GetNation())}, nil
}

func (game *Game) ViewUnits(playerId int, request api.ViewUnitsRequest) (*api.ViewUnitsResponse, error) {
	units := make(map[string]*api.UnitStruct, 0)
	for loc, unit := range game.world.GetUnits() {
		units[fmt.Sprintf("%d", unit.GetId())] = unit.Serialize(loc)
	}
	return &api.ViewUnitsResponse{
		Units: units}, nil
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
