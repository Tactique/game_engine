package engine

import (
    "encoding/json"
    "requests"
    "engine/game_engine"
    "fmt"
)

func respondSuccess(payload interface{}) string {
    return generateResponse(payload, 0)
}

func respondFailure(payload interface{}) string {
    return generateResponse(payload, 1)
}

func respondMalformed(payload interface{}) string {
    return generateResponse(payload, 2)
}

func respondBadCommand(payload interface{}) string {
    return generateResponse(payload, 3)
}

func respondUnknownCommand(payload interface{}) string {
    return generateResponse(payload, 4)
}

func generateResponse(payload interface{}, status int) string {
    response, err := json.Marshal(map[string]interface{}{"status": status, "payload": payload})
    if err != nil {
        fmt.Println(err.Error())
        return "{\"status\": 5, \"payload\": \"oops\"}"
    }
    return string(response)
}

func newCommand(jsonRequest string) (string, *game_engine.Game) {
    var request requests.NewCommandRequest
    err := json.Unmarshal([]byte(jsonRequest), &request)
    if err != nil {
        return respondMalformed(nil), nil
    }
    game, err := game_engine.NewGame(request.Uids, request.Debug)
    if err != nil {
        return respondBadCommand(err.Error()), nil
    } else {
        return respondSuccess(nil), game
    }
}


func viewCommand(jsonRequest string, playerId int, game *game_engine.Game) string {
    var request requests.ViewCommandRequest
    err := json.Unmarshal([]byte(jsonRequest), &request)
    if err != nil {
        return respondMalformed(nil)
    }
    worldStruct, err := game.Serialize(playerId)
    if err != nil {
        return respondBadCommand(err.Error())
    }
    return respondSuccess(worldStruct)
}

func moveCommand(jsonRequest string, playerId int, game *game_engine.Game) string {
    var request requests.MoveCommandRequest
    err := json.Unmarshal([]byte(jsonRequest), &request)
    if err != nil {
        return respondMalformed(nil)
    }
    moveErr := game.MoveUnit(playerId, request.Move)
    if moveErr != nil {
        return respondBadCommand(moveErr.Error())
    }
    return respondSuccess(nil)
}

func attackCommand(jsonRequest string, playerId int, game *game_engine.Game) string {
    var request requests.AttackCommandRequest
    err := json.Unmarshal([]byte(jsonRequest), &request)
    if err != nil {
        return respondMalformed(nil)
    }
    attackErr := game.Attack(playerId, request.Attacker, request.AttackIndex, request.Target)
    if attackErr != nil {
        return respondBadCommand(attackErr.Error())
    }
    return respondSuccess(nil)
}

func endTurnCommand(jsonRequest string, playerId int, game *game_engine.Game) string {
    var request requests.MoveCommandRequest
    err := json.Unmarshal([]byte(jsonRequest), &request)
    if err != nil {
        return respondMalformed(nil)
    }
    endErr := game.EndTurn(playerId)
    if endErr != nil {
        return respondBadCommand(endErr.Error())
    }
    return respondSuccess(nil)
}

func exitCommand(jsonRequest string, playerId int, game *game_engine.Game) string {
    var request requests.ExitCommandRequest
    err := json.Unmarshal([]byte(jsonRequest), &request)
    if err != nil {
        return respondMalformed(nil)
    }
    return respondSuccess(nil)
}
