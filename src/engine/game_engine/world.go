package game_engine

import (
	"api"
	"errors"
	"fmt"
	"github.com/Tactique/golib/logger"
)

type World struct {
	terrain    [][]terrain
	unitMap    map[location]*unit
	players    []*player
	numPlayers int
	turnOwner  int
	nextUnitId int
}

func NewWorld(playerIds []int, worldId int) (*World, error) {
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

	if len(terrains) < 2 {
		return nil, errors.New("No terrains were loadable")
	}
	plains := terrains[0]
	roads := terrains[1]

	ret_world := &World{
		terrain: [][]terrain{
			[]terrain{plains, roads, plains, plains, plains, plains, plains, plains},
			[]terrain{plains, roads, plains, plains, plains, plains, plains, plains},
			[]terrain{roads, roads, roads, plains, plains, plains, plains, plains},
			[]terrain{plains, plains, roads, plains, plains, plains, plains, plains},
			[]terrain{plains, plains, roads, plains, plains, plains, plains, plains},
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
		ret_world.AddUnit(newLocation(0, 0), name, nations[0], dbHealth, dbAttacks, dbArmor, dbMovement)
		name = "mage"
		dbHealth, dbAttacks, dbArmor, dbMovement, err = loadUnit(db, name)
		if err != nil {
			return nil, err
		}
		ret_world.AddUnit(newLocation(3, 3), name, nations[0], dbHealth, dbAttacks, dbArmor, dbMovement)
		if numPlayers == 2 {
			ret_world.AddUnit(newLocation(0, 3), name, nations[1], dbHealth, dbAttacks, dbArmor, dbMovement)
		}
	}
	return ret_world, nil
}

func (world *World) getPlayer(playerId int) (*player, error) {
	for _, player := range world.players {
		if player.playerId == playerId {
			return player, nil
		}
	}
	return nil, errors.New("Player not playing")
}

func (world *World) verifyTurnOwner(playerId int) error {
	if playerId != world.players[world.turnOwner].playerId {
		return errors.New("Not the turn owner")
	}
	return nil
}

func (world *World) getAndVerifyTurnOwner(playerId int) (*player, error) {
	err := world.verifyTurnOwner(playerId)
	if err != nil {
		return nil, err
	}
	return world.getPlayer(playerId)
}

func (world *World) getUnit(location location) (*unit, error) {
	unit, ok := world.unitMap[location]
	if ok {
		return unit, nil
	} else {
		message := fmt.Sprintf("No unit located at (%d, %d)", location.x, location.y)
		logger.Warn(message)
		return nil, errors.New(message)
	}

}

func (world *World) verifyOwnedUnit(player *player, unit *unit) error {
	if unit.nation != player.nation {
		logger.Warnf("Unit owned by %d is not owned by the current player (%d)", unit.nation, player.nation)
		return errors.New("Unit is not owned by the current player")
	} else {
		return nil
	}
}

func (world *World) getAndVerifyOwnedUnit(player *player, location location) (*unit, error) {
	unit, err := world.getUnit(location)
	if err != nil {
		return nil, err
	}
	return unit, world.verifyOwnedUnit(player, unit)
}

func (world *World) AddUnit(
	location location, name string, nation nation,
	health int, attacks []*attack, armor *armor, movement *movement) error {
	logger.Infof("Adding unit at (x: %d, y: %d)", location.x, location.y)
	_, ok := world.unitMap[location]
	if !ok {
		world.unitMap[location] = newUnit(
			name, world.nextUnitId, nation, health, attacks, armor, movement)
		world.nextUnitId += 1
		logger.Infof("Added unit at (x: %d, y: %d)", location.x, location.y)
		return nil
	} else {
		logger.Warnf("Failed to add unit at (x: %d, y: %d)", location.x, location.y)
		return errors.New("location already occupied")
	}
}

func (world *World) GetUnits() map[location]unit {
	copiedUnitMap := make(map[location]unit, len(world.unitMap))
	for loc, unit := range world.unitMap {
		copiedUnitMap[loc] = *unit
	}
	return copiedUnitMap
}

func (world *World) GetTerrain() [][]terrain {
	return world.terrain
}

func (world *World) GetPlayers() []player {
	players := make([]player, world.numPlayers)
	for i, player := range world.players {
		players[i] = *player
	}
	return players
}

func (world *World) GetNumPlayers() int {
	return world.numPlayers
}

func (world *World) GetTurnOwner() *player {
	return world.players[world.turnOwner]
}

func (world *World) MoveUnit(playerId int, rawLocations []*api.LocationStruct) (*api.MoveResponse, error) {
	player, err := world.getAndVerifyTurnOwner(playerId)
	if err != nil {
		return nil, err
	}
	locations := make([]location, len(rawLocations))
	for i, location := range rawLocations {
		locations[i] = locationFromRequest(location)
	}
	validError := world.verifyValidMove(player, locations)
	if validError != nil {
		return nil, validError
	}
	world.verifiedMoveUnit(locations)
	return &api.MoveResponse{
		Move: rawLocations}, nil
}

func (world *World) verifyValidMove(player *player, locations []location) error {
	if len(locations) < 1 {
		message := "must supply more than zero locations"
		logger.Warnf(message)
		return errors.New(message)
	}

	tiles := make([]terrain, len(locations))
	for i, location := range locations {
		tiles[i] = world.terrain[location.x][location.y]
	}
	for _, location := range locations[1:] {
		if world.unitMap[location] != nil {
			return errors.New("Cannot pass through units")
		}
	}
	unit, err := world.getAndVerifyOwnedUnit(player, locations[0])
	if err != nil {
		return err
	}
	return validMove(unit.movement.distance, unit.movement, tiles, locations)
}

func (world *World) verifiedMoveUnit(locations []location) error {
	end := len(locations)
	unit := world.unitMap[newLocation(locations[0].x, locations[0].y)]
	unit.canMove = false
	world.unitMap[newLocation(locations[end-1].x, locations[end-1].y)] = unit
	delete(world.unitMap, newLocation(locations[0].x, locations[0].y))
	return nil
}

func (world *World) Attack(
	playerId int, attacker *api.LocationStruct,
	attackIndex int, target *api.LocationStruct) (*api.AttackResponse, error) {
	player, err := world.getAndVerifyTurnOwner(playerId)
	if err != nil {
		return nil, err
	}
	attackingUnit, err := world.getAndVerifyOwnedUnit(player, locationFromRequest(attacker))
	if err != nil {
		return nil, err
	}
	defendingUnit, err := world.getUnit(locationFromRequest(target))
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

func (world *World) EndTurn(playerId int) (*api.EndTurnResponse, error) {
	err := world.verifyTurnOwner(playerId)
	if err != nil {
		return nil, err
	}
	nextOwner := world.turnOwner + 1
	if nextOwner >= world.numPlayers {
		world.turnOwner = 0
	} else {
		world.turnOwner = nextOwner
	}
	currentOwner := world.players[world.turnOwner]
	units := make(map[string]*api.UnitStruct, 0)
	for loc, unit := range world.unitMap {
		if unit.nation == currentOwner.nation {
			unit.turnReset()
			units[fmt.Sprintf("%d", unit.id)] = unit.Serialize(loc)
		}
	}
	return &api.EndTurnResponse{
		PlayerId:     playerId,
		ChangedUnits: units}, nil
}
