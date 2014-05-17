package request_handler

import (
	"api"
	"engine/game"
	"errors"
	"fmt"
	"github.com/Tactique/golib/logger"
)

type GameWrapper struct {
	world *game.World
}

func NewGameWrapper(request api.NewRequest) (*GameWrapper, error) {
	world, err := game.NewWorld(request.Uids, request.Debug)
	return &GameWrapper{world: world}, err
}

func (gameWrapper *GameWrapper) ViewWorld(playerId int, request api.ViewWorldRequest) (*api.ViewWorldResponse, error) {
	terrain, err := gameWrapper.ViewTerrain(playerId, api.ViewTerrainRequest{})
	if err != nil {
		return nil, err
	}
	players, err := gameWrapper.ViewPlayers(playerId, api.ViewPlayersRequest{})
	if err != nil {
		return nil, err
	}
	units, err := gameWrapper.ViewUnits(playerId, api.ViewUnitsRequest{})
	if err != nil {
		return nil, err
	}
	return &api.ViewWorldResponse{
		TerrainResponse: terrain,
		UnitsResponse:   units,
		PlayersResponse: players}, nil
}

func (gameWrapper *GameWrapper) ViewTerrain(playerId int, request api.ViewTerrainRequest) (*api.ViewTerrainResponse, error) {
	terrain := gameWrapper.world.GetTerrain()
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

func (gameWrapper *GameWrapper) ViewPlayers(playerId int, request api.ViewPlayersRequest) (*api.ViewPlayersResponse, error) {
	players := make(map[string]*api.PlayerStruct, gameWrapper.world.GetNumPlayers())
	for _, player := range gameWrapper.world.GetPlayers() {
		players[fmt.Sprintf("%d", player.GetPlayerId())] = player.Serialize()
	}
	return &api.ViewPlayersResponse{
		Players:   players,
		TurnOwner: int(gameWrapper.world.GetTurnOwner().GetNation())}, nil
}

func (gameWrapper *GameWrapper) ViewUnits(playerId int, request api.ViewUnitsRequest) (*api.ViewUnitsResponse, error) {
	units := make(map[string]*api.UnitStruct, 0)
	for loc, unit := range gameWrapper.world.GetUnits() {
		units[fmt.Sprintf("%d", unit.GetId())] = unit.Serialize(loc)
	}
	return &api.ViewUnitsResponse{
		Units: units}, nil
}

func (gameWrapper *GameWrapper) MoveUnit(playerId int, request api.MoveRequest) (*api.MoveResponse, error) {
	player, err := gameWrapper.world.GetAndVerifyTurnOwner(playerId)
	if err != nil {
		return nil, err
	}
	locations := make([]game.Location, len(request.Move))
	for i, location := range request.Move {
		locations[i] = game.LocationFromRequest(location)
	}
	validError := gameWrapper.verifyValidMove(request.UnitId, player, locations)
	if validError != nil {
		return nil, validError
	}
	end := len(locations)
	err = gameWrapper.verifiedMoveUnit(request.UnitId, locations[0], locations[end-1])
	if err != nil {
		return nil, err
	}
	return &api.MoveResponse{
		Move:   request.Move,
		UnitId: request.UnitId}, nil
}

func (gameWrapper *GameWrapper) verifyValidMove(unitId int, player *game.Player, locations []game.Location) error {
	if len(locations) < 1 {
		message := "must supply more than zero locations"
		logger.Warnf(message)
		return errors.New(message)
	}

	tiles := make([]game.Terrain, len(locations))
	terrain := gameWrapper.world.GetTerrain()
	for i, location := range locations {
		tiles[i] = terrain[location.GetX()][location.GetY()]
	}
	for _, location := range locations[1:] {
		err := gameWrapper.world.IsUnitAtLocation(location)
		if err != nil {
			logger.Warn(err)
			return errors.New(fmt.Sprintf("Cannot pass through units at (%d, %d)", location.GetX(), location.GetY()))
		}
	}
	unit, err := gameWrapper.world.GetAndVerifyOwnedUnit(player, locations[0])
	if err != nil {
		return err
	}
	return game.ValidMove(unit.GetMovement().GetDistance(), unit.GetMovement(), tiles, locations)
}

func (gameWrapper *GameWrapper) verifiedMoveUnit(unitId int, start game.Location, end game.Location) error {
	return gameWrapper.world.MoveUnitFromTo(unitId, start, end)
}

func (gameWrapper *GameWrapper) Attack(playerId int, request api.AttackRequest) (*api.AttackResponse, error) {
	return gameWrapper.world.Attack(playerId, request.Attacker, request.AttackIndex, request.Target)
}

func (gameWrapper *GameWrapper) EndTurn(playerId int, request api.EndTurnRequest) (*api.EndTurnResponse, error) {
	return gameWrapper.world.EndTurn(playerId)
}
