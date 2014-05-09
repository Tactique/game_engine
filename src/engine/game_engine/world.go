package game_engine

import (
	"api"
	"errors"
	"fmt"
	"github.com/Tactique/golib/logger"
)

type Game struct {
	terrain    [][]terrain
	unitMap    map[location]*unit
	players    []*player
	numPlayers int
	turnOwner  int
	nextUnitId int
}

func NewGame(playerIds []int, worldId int) (*Game, error) {
	db, err := newDatabase()
	if err != nil {
		logger.Errorf("DB is open and cannot be used (%s)", err.Error())
		return nil, err
	}
	defer db.Close()
	terrains, err := loadTerrains(db)
	if err != nil {
		return nil, err
	}
	nations, err := loadNations(db)
	if err != nil {
		return nil, err
	}
	numPlayers := len(playerIds)
	if numPlayers > 4 || numPlayers < 1 {
		logger.Warnf("Must have between 1 and 4 players, got %d", numPlayers)
		return nil, errors.New("must have between 1 and 4 players")
	}
	if len(nations) < numPlayers {
		logger.Errorf("Not enough nations were loaded, must have at least 2, got %s", nations)
		return nil, errors.New("Not enough nations were loaded, must have at least 2")
	}
	players := make([]*player, numPlayers)
	for i, playerId := range playerIds {
		players[i] = newPlayer(playerId, nations[i], team(i))
	}

	if len(terrains) < 1 {
		return nil, errors.New("No terrains were loadable")
	}
	plains := terrains[0]

	ret_game := &Game{
		terrain: [][]terrain{
			[]terrain{plains, plains, plains, plains, plains, plains, plains, plains},
			[]terrain{plains, plains, plains, plains, plains, plains, plains, plains},
			[]terrain{plains, plains, plains, plains, plains, plains, plains, plains},
			[]terrain{plains, plains, plains, plains, plains, plains, plains, plains},
			[]terrain{plains, plains, plains, plains, plains, plains, plains, plains},
			[]terrain{plains, plains, plains, plains, plains, plains, plains, plains}},
		unitMap:    make(map[location]*unit),
		players:    players,
		numPlayers: numPlayers,
		turnOwner:  0,
		nextUnitId: 0}
	if worldId == 0 {
		name := "warrior"
		dbHealth, dbAttacks, dbArmor, dbMovement, err := loadUnit(db, name)
		if err != nil {
			return nil, err
		}
		ret_game.AddUnit(newLocation(0, 0), name, nations[0], dbHealth, dbAttacks, dbArmor, dbMovement)
		name = "mage"
		dbHealth, dbAttacks, dbArmor, dbMovement, err = loadUnit(db, name)
		if err != nil {
			return nil, err
		}
		ret_game.AddUnit(newLocation(3, 3), name, nations[0], dbHealth, dbAttacks, dbArmor, dbMovement)
		if numPlayers == 2 {
			ret_game.AddUnit(newLocation(0, 3), name, nations[1], dbHealth, dbAttacks, dbArmor, dbMovement)
		}
	}
	return ret_game, nil
}

func (game *Game) getPlayer(playerId int) (*player, error) {
	for _, player := range game.players {
		if player.playerId == playerId {
			return player, nil
		}
	}
	return nil, errors.New("Player not playing")
}

func (game *Game) verifyTurnOwner(playerId int) error {
	if playerId != game.players[game.turnOwner].playerId {
		return errors.New("Not the turn owner")
	}
	return nil
}

func (game *Game) getAndVerifyTurnOwner(playerId int) (*player, error) {
	err := game.verifyTurnOwner(playerId)
	if err != nil {
		return nil, err
	}
	return game.getPlayer(playerId)
}

func (game *Game) getUnit(location location) (*unit, error) {
	unit, ok := game.unitMap[location]
	if ok {
		return unit, nil
	} else {
		message := fmt.Sprintf("No unit located at (%d, %d)", location.x, location.y)
		logger.Warn(message)
		return nil, errors.New(message)
	}

}

func (game *Game) verifyOwnedUnit(player *player, unit *unit) error {
	if unit.nation != player.nation {
		logger.Warnf("Unit owned by %d is not owned by the current player (%d)", unit.nation, player.nation)
		return errors.New("Unit is not owned by the current player")
	} else {
		return nil
	}
}

func (game *Game) getAndVerifyOwnedUnit(player *player, location location) (*unit, error) {
	unit, err := game.getUnit(location)
	if err != nil {
		return nil, err
	}
	return unit, game.verifyOwnedUnit(player, unit)
}

func (game *Game) AddUnit(
	location location, name string, nation nation,
	health int, attacks []*attack, armor *armor, movement *movement) error {
	logger.Infof("Adding unit at (x: %d, y: %d)", location.x, location.y)
	_, ok := game.unitMap[location]
	if !ok {
		game.unitMap[location] = newUnit(
			name, game.nextUnitId, nation, health, attacks, armor, movement)
		game.nextUnitId += 1
		logger.Infof("Added unit at (x: %d, y: %d)", location.x, location.y)
		return nil
	} else {
		logger.Warnf("Failed to add unit at (x: %d, y: %d)", location.x, location.y)
		return errors.New("location already occupied")
	}
}

func (game *Game) ViewWorld(playerId int) (*api.ViewWorldResponse, error) {
	terrain, err := game.ViewTerrain(playerId)
	if err != nil {
		return nil, err
	}
	players, err := game.ViewPlayers(playerId)
	if err != nil {
		return nil, err
	}
	units, err := game.ViewUnits(playerId)
	if err != nil {
		return nil, err
	}
	return &api.ViewWorldResponse{
		TerrainResponse: terrain,
		UnitsResponse:   units,
		PlayersResponse: players}, nil
}

func (game *Game) ViewTerrain(playerId int) (*api.ViewTerrainResponse, error) {
	terrainInts := make([][]int, len(game.terrain))
	for i, t := range game.terrain {
		thoriz := make([]int, len(t))
		for j, t_ := range t {
			thoriz[j] = int(t_)
		}
		terrainInts[i] = thoriz
	}
	return &api.ViewTerrainResponse{
		Terrain: terrainInts}, nil
}

func (game *Game) ViewUnits(playerid int) (*api.ViewUnitsResponse, error) {
	units := make(map[string]*api.UnitStruct, 0)
	for loc, unit := range game.unitMap {
		units[fmt.Sprintf("%d", unit.id)] = unit.serialize(loc)
	}
	return &api.ViewUnitsResponse{
		Units: units}, nil

}

func (game *Game) ViewPlayers(playerId int) (*api.ViewPlayersResponse, error) {
	players := make(map[string]*api.PlayerStruct, game.numPlayers)
	for _, player := range game.players {
		players[fmt.Sprintf("%d", player.playerId)] = player.serialize()
	}
	return &api.ViewPlayersResponse{
		Players:   players,
		TurnOwner: int(game.players[game.turnOwner].nation)}, nil
}

func (game *Game) MoveUnit(playerId int, rawLocations []*api.LocationStruct) (*api.MoveResponse, error) {
	player, err := game.getAndVerifyTurnOwner(playerId)
	if err != nil {
		return nil, err
	}
	locations := make([]location, len(rawLocations))
	for i, location := range rawLocations {
		locations[i] = locationFromRequest(location)
	}
	validError := game.verifyValidMove(player, locations)
	if validError != nil {
		return nil, validError
	}
	game.verifiedMoveUnit(locations)
	return &api.MoveResponse{
		Move: rawLocations}, nil
}

func (game *Game) verifyValidMove(player *player, locations []location) error {
	if len(locations) < 1 {
		message := "must supply more than zero locations"
		logger.Warnf(message)
		return errors.New(message)
	}

	tiles := make([]terrain, len(locations))
	for i, location := range locations {
		tiles[i] = game.terrain[location.x][location.y]
	}
	for _, location := range locations[1:] {
		if game.unitMap[location] != nil {
			return errors.New("Cannot pass through units")
		}
	}
	unit, err := game.getAndVerifyOwnedUnit(player, locations[0])
	if err != nil {
		return err
	}
	return validMove(unit.movement.distance, unit.movement, tiles, locations)
}

func (game *Game) verifiedMoveUnit(locations []location) error {
	end := len(locations)
	unit := game.unitMap[newLocation(locations[0].x, locations[0].y)]
	unit.canMove = false
	game.unitMap[newLocation(locations[end-1].x, locations[end-1].y)] = unit
	delete(game.unitMap, newLocation(locations[0].x, locations[0].y))
	return nil
}

func (game *Game) Attack(
	playerId int, attacker *api.LocationStruct,
	attackIndex int, target *api.LocationStruct) (*api.AttackResponse, error) {
	player, err := game.getAndVerifyTurnOwner(playerId)
	if err != nil {
		return nil, err
	}
	attackingUnit, err := game.getAndVerifyOwnedUnit(player, locationFromRequest(attacker))
	if err != nil {
		return nil, err
	}
	defendingUnit, err := game.getUnit(locationFromRequest(target))
	if err != nil {
		return nil, err
	}
	alive, err := damageUnit(attackingUnit, attackIndex, defendingUnit)
	// TODO When a unit is dead mark is as such
	logger.Errorf("unit is %s (alive)", alive)
	if err != nil {
		return nil, err
	}
	return &api.AttackResponse{
		Attacker: attacker, AttackIndex: attackIndex,
		Target: target}, nil
}

func (game *Game) EndTurn(playerId int) (*api.EndTurnResponse, error) {
	err := game.verifyTurnOwner(playerId)
	if err != nil {
		return nil, err
	}
	nextOwner := game.turnOwner + 1
	if nextOwner >= game.numPlayers {
		game.turnOwner = 0
	} else {
		game.turnOwner = nextOwner
	}
	currentOwner := game.players[game.turnOwner]
	units := make(map[string]*api.UnitStruct, 0)
	for loc, unit := range game.unitMap {
		if unit.nation == currentOwner.nation {
			unit.turnReset()
			units[fmt.Sprintf("%d", unit.id)] = unit.serialize(loc)
		}
	}
	return &api.EndTurnResponse{
		PlayerId:     playerId,
		ChangedUnits: units}, nil
}
