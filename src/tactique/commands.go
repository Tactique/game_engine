package tactique

import (
    "encoding/json"
    "requests"
    "tactique/game_engine"
)

func handleCommand(jsonRequest string, request interface{}) (string, error) {
    return "", nil
}

func newCommand(jsonRequest string) (string, *game_engine.Game, error) {
    var request requests.NewCommandRequest
    err := json.Unmarshal([]byte(jsonRequest), &request)
    if err != nil {
        return "", nil, err
    }
    game, err := game_engine.NewGame(request.Uids, 1)
    if err != nil {
        return "new:failure:" + err.Error(), nil, nil
    } else {
        return "new:success", game, nil
    }
}

func viewCommand(jsonRequest string, playerId int, game *game_engine.Game) (string, error) {
    var request requests.ViewCommandRequest
    err := json.Unmarshal([]byte(jsonRequest), &request)
    if err != nil {
        return "", err
    }
    bytestring, err := game.Serialize(playerId)
    if err != nil {
        return "", err
    }
    return "view:success:" + string(bytestring), nil
}

func moveCommand(jsonRequest string, playerId int, game *game_engine.Game) (string, error) {
    var request requests.MoveCommandRequest
    err := json.Unmarshal([]byte(jsonRequest), &request)
    if err != nil {
        return "", err
    }
    moveErr := game.MoveUnit(playerId, request.Move)
    if moveErr != nil {
        return "move:failure:" + moveErr.Error(), nil
    } else {
        return "move:success", nil
    }
}

func endTurnCommand(jsonRequest string, playerId int, game *game_engine.Game) (string, error) {
    var request requests.MoveCommandRequest
    err := json.Unmarshal([]byte(jsonRequest), &request)
    if err != nil {
        return "", err
    }
    moveErr := game.EndTurn(playerId)
    if moveErr != nil {
        return "end:failure:" + moveErr.Error(), nil
    } else {
        return "end:success", nil
    }
}

func exitCommand(jsonRequest string, playerId int, game *game_engine.Game) (string, error) {
    var request requests.ExitCommandRequest
    err := json.Unmarshal([]byte(jsonRequest), &request)
    if err != nil {
        return "", err
    }
    return "exit", nil
}
