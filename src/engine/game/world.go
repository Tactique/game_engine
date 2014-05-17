package game

import (
	"api"
	"errors"
	"fmt"
	"github.com/Tactique/golib/logger"
)

type World struct {
	terrain    [][]Terrain
	unitMap    map[Location]*Unit
	units      map[int]*Unit
	players    []*Player
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
	players := make([]*Player, numPlayers)
	for i, playerId := range playerIds {
		players[i] = NewPlayer(playerId, nations[i], team(i))
	}

	if len(terrains) < 2 {
		return nil, errors.New("No terrains were loadable")
	}
	plains := terrains[0]
	roads := terrains[1]

	ret_world := &World{
		terrain: [][]Terrain{
			[]Terrain{plains, roads, plains, plains, plains, plains, plains, plains},
			[]Terrain{plains, roads, plains, plains, plains, plains, plains, plains},
			[]Terrain{roads, roads, roads, plains, plains, plains, plains, plains},
			[]Terrain{plains, plains, roads, plains, plains, plains, plains, plains},
			[]Terrain{plains, plains, roads, plains, plains, plains, plains, plains},
			[]Terrain{plains, plains, plains, plains, plains, plains, plains, plains}},
		unitMap:    make(map[Location]*Unit),
		units:      make(map[int]*Unit),
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
		ret_world.AddUnit(NewLocation(0, 0), name, nations[0], dbHealth, dbAttacks, dbArmor, dbMovement)
		name = "mage"
		dbHealth, dbAttacks, dbArmor, dbMovement, err = loadUnit(db, name)
		if err != nil {
			return nil, err
		}
		ret_world.AddUnit(NewLocation(3, 3), name, nations[0], dbHealth, dbAttacks, dbArmor, dbMovement)
		if numPlayers == 2 {
			ret_world.AddUnit(NewLocation(0, 3), name, nations[1], dbHealth, dbAttacks, dbArmor, dbMovement)
		}
	}
	return ret_world, nil
}

func (world *World) getPlayer(playerId int) (*Player, error) {
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

func (world *World) GetAndVerifyTurnOwner(playerId int) (*Player, error) {
	err := world.verifyTurnOwner(playerId)
	if err != nil {
		return nil, err
	}
	return world.getPlayer(playerId)
}

func (world *World) IsUnitAtLocation(location Location) error {
	_, ok := world.unitMap[location]
	if ok {
		logger.Warn(world.unitMap)
		return errors.New(fmt.Sprintf("There is a unit at the  location (%d, %d)", location.x, location.y))
	} else {
		return nil
	}
}

func (world *World) GetUnitFromId(unitId int) (*Unit, error) {
	unit, ok := world.units[unitId]
	if ok {
		return unit, nil
	} else {
		message := fmt.Sprintf("No unit with id %d", unitId)
		logger.Warn(message)
		return nil, errors.New(message)
	}
}

func (world *World) GetUnitAtLocation(location Location) (*Unit, error) {
	unit, ok := world.unitMap[location]
	if ok {
		return unit, nil
	} else {
		message := fmt.Sprintf("No unit located at (%d, %d)", location.x, location.y)
		logger.Warn(message)
		return nil, errors.New(message)
	}
}

func (world *World) verifyOwnedUnit(player *Player, unit *Unit) error {
	if unit.nation != player.nation {
		logger.Warnf("Unit owned by %d is not owned by the current player (%d)", unit.nation, player.nation)
		return errors.New("Unit is not owned by the current player")
	} else {
		return nil
	}
}

func (world *World) GetAndVerifyOwnedUnit(player *Player, unitId int) (*Unit, error) {
	unit, err := world.GetUnitFromId(unitId)
	if err != nil {
		return nil, err
	}
	return unit, world.verifyOwnedUnit(player, unit)
}

func (world *World) AddUnit(
	location Location, name string, nation nation,
	health int, attacks []*attack, armor *armor, movement *Movement) error {
	logger.Infof("Adding unit at (x: %d, y: %d)", location.x, location.y)
	unitId := world.nextUnitId
	_, okUnitLocation := world.unitMap[location]
	_, okUnitId := world.units[unitId]
	if !(okUnitLocation && okUnitId) {
		newUnit := NewUnit(
			name, unitId, nation, health, attacks, armor, movement)
		world.unitMap[location] = newUnit
		world.units[unitId] = newUnit
		world.nextUnitId += 1
		logger.Infof("Added unit with id %d at (x: %d, y: %d)", unitId, location.x, location.y)
		return nil
	} else {
		logger.Warnf("Failed to add unit with id %d at (x: %d, y: %d)", unitId, location.x, location.y)
		return errors.New("location already occupied")
	}
}

func (world *World) GetUnits() map[Location]Unit {
	copiedUnitMap := make(map[Location]Unit, len(world.unitMap))
	for loc, unit := range world.unitMap {
		copiedUnitMap[loc] = *unit
	}
	return copiedUnitMap
}

func (world *World) GetTerrain() [][]Terrain {
	return world.terrain
}

func (world *World) GetPlayers() []Player {
	players := make([]Player, world.numPlayers)
	for i, player := range world.players {
		players[i] = *player
	}
	return players
}

func (world *World) GetNumPlayers() int {
	return world.numPlayers
}

func (world *World) GetTurnOwner() *Player {
	return world.players[world.turnOwner]
}

func (world *World) MoveUnitFromTo(unitId int, start Location, end Location) error {
	unit, err := world.GetUnitAtLocation(start)
	if err != nil {
		return err
	}
	verifiedUnit, err := world.GetUnitFromId(unitId)
	if err != nil {
		return err
	}
	if unit != verifiedUnit {
		message := fmt.Sprintf("Unit with given unitId (%d) isn't at the specified location (%d, %d)", unitId, start.x, start.y)
		logger.Warnf(message)
		return errors.New(message)
	}
	unit.canMove = false
	err = world.IsUnitAtLocation(end)
	if err != nil {
		return err
	} else {
		world.unitMap[end] = unit
		delete(world.unitMap, start)
		return nil
	}
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
