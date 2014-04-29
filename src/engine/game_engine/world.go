package game_engine

import (
    "fmt"
    "encoding/json"
    "errors"
    "requests"
)

type Game struct {
    terrain [][]terrain
    unitMap map[location]*unit
    players []*player
    numPlayers int
    turnOwner int
}

func NewGame(playerIds []int, worldId int) (*Game, error) {
    numPlayers := len(playerIds)
    if numPlayers > 4 || numPlayers < 1 {
        return nil, errors.New("must have between 1 and 4 players")
    }
    players := make([]*player, numPlayers)
    for i, playerId := range(playerIds) {
        players[i] = newPlayer(playerId, nation(i), team(i))
    }
    ret_game := &Game{
        terrain: [][]terrain{
            []terrain{plains, plains, plains, plains, plains, plains, plains, plains},
            []terrain{plains, plains, plains, plains, plains, plains, plains, plains},
            []terrain{plains, plains, plains, plains, plains, plains, plains, plains},
            []terrain{plains, plains, plains, plains, plains, plains, plains, plains},
            []terrain{plains, plains, plains, plains, plains, plains, plains, plains},
            []terrain{plains, plains, plains, plains, plains, plains, plains, plains}},
        unitMap: make(map[location]*unit),
        players: players,
        numPlayers: numPlayers,
        turnOwner: 0}
    if worldId == 0 {
        ret_game.AddUnit(newLocation(0, 0), warrior(red))
        ret_game.AddUnit(newLocation(0, 3), warrior(blue))
    }
    return ret_game, nil
}


func (game *Game) getPlayer(playerId int) (*player, error) {
    for _, player := range(game.players) {
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
    unit, ok := game.unitMap[location]; if ok {
        return unit, nil
    } else {
        message := fmt.Sprintf("No unit located at (%d, %d)", location.x, location.y)
        fmt.Println(message)
        return nil, errors.New(message)
    }

}

func (game *Game) verifyOwnedUnit(player *player, unit *unit) error {
    if unit.nation != player.nation {
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

func (game *Game) AddUnit(location location, unit *unit) error {
    fmt.Println("adding unit")
    _, ok := game.unitMap[location]; if !ok {
        game.unitMap[location] = unit
        fmt.Println("added unit")
        return nil
    } else {
        fmt.Println("failed to add unit")
        return errors.New("location already occupied")
    }
}

func (game *Game) Serialize(playerId int) ([]byte, error) {
    players := make([]*requests.PlayerStruct, len(game.players))
    for i, player := range(game.players) {
        players[i] = player.serialize()
    }
    terrainInts := make([][]int, len(game.terrain))
    for i, t := range(game.terrain) {
        thoriz := make([]int, len(t))
        for j, t_ := range(t) {
            thoriz[j] = int(t_)
        }
        terrainInts[i] = thoriz
    }
    units := make([]*requests.UnitStruct, len(game.unitMap))
    i := 0
    for location, unit := range(game.unitMap) {
        units[i] = unit.serialize(location)
        i += 1
    }
    return json.Marshal(requests.WorldStruct{
        Terrain: terrainInts,
        Units: units,
        Players: players,
        TurnOwner: game.players[game.turnOwner].playerId})
}

func (game *Game) MoveUnit(playerId int, rawLocations []requests.LocationStruct) error {
    player, err := game.getAndVerifyTurnOwner(playerId)
    if err != nil {
        return err
    }
    locations := make([]location, len(rawLocations))
    for i, location := range(rawLocations) {
        locations[i] = locationFromRequest(location)
    }
    validError := game.verifyValidMove(player, locations)
    if validError != nil {
        return validError
    }
    game.verifiedMoveUnit(locations)
    return nil
}

func (game *Game) verifyValidMove(player *player, locations []location) error {
    if len(locations) < 1 {
        message := "must supply more than zero locations"
        fmt.Println(message)
        return errors.New(message)
    }

    tiles := make([]terrain, len(locations))
    for i, location := range(locations) {
        tiles[i] = game.terrain[location.x][location.y]
    }
    for _, location := range(locations[1:]) {
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
        playerId int, attacker requests.LocationStruct,
        attackIndex int, target requests.LocationStruct) error {
    return nil
}

func (game *Game) EndTurn(playerId int) error {
    player, err := game.getAndVerifyTurnOwner(playerId)
    if err != nil {
        return err
    }
    for _, unit := range(game.unitMap) {
        if unit.nation == player.nation {
            unit.turnReset()
        }
    }
    nextOwner := game.turnOwner + 1
    if nextOwner >= game.numPlayers {
        game.turnOwner = 0
    } else {
        game.turnOwner = nextOwner
    }
    return nil
}
