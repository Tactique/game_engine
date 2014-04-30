package engine

import (
    "encoding/json"
    "requests"
    "engine/game_engine"
    "fmt"
)

func respondSuccess(payload interface{}) []byte {
    return generateResponse(payload, 0)
}

func respondFailure(payload interface{}) []byte {
    return generateResponse(payload, 1)
}

func respondMalformed(payload interface{}) []byte {
    return generateResponse(payload, 2)
}

func respondBadCommand(payload interface{}) []byte {
    return generateResponse(payload, 3)
}

func respondUnknownCommand(payload interface{}) []byte {
    return generateResponse(payload, 4)
}

func generateResponse(payload interface{}, status int) []byte {
    response, err := json.Marshal(map[string]interface{}{"status": status, "payload": payload})
    if err != nil {
        fmt.Println(err.Error())
        return []byte("{\"status\": 5, \"payload\": \"oops\"}")
    }
    return response
}

func newCommand(jsonRequest []byte) ([]byte, *game_engine.Game) {
    var request requests.NewCommandRequest
    err := json.Unmarshal(jsonRequest, &request)
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


func viewCommand(jsonRequest []byte, playerId int, game *game_engine.Game) []byte {
    var request requests.ViewCommandRequest
    err := json.Unmarshal(jsonRequest, &request)
    if err != nil {
        return respondMalformed(nil)
    }
    worldStruct, err := game.Serialize(playerId)
    if err != nil {
        return respondBadCommand(err.Error())
    }
    return respondSuccess(worldStruct)
}

func moveCommand(jsonRequest []byte, playerId int, game *game_engine.Game) []byte {
    var request requests.MoveCommandRequest
    err := json.Unmarshal(jsonRequest, &request)
    if err != nil {
        return respondMalformed(nil)
    }
    moveErr := game.MoveUnit(playerId, request.Move)
    if moveErr != nil {
        return respondBadCommand(moveErr.Error())
    }
    return respondSuccess(nil)
}

func attackCommand(jsonRequest []byte, playerId int, game *game_engine.Game) []byte {
    var request requests.AttackCommandRequest
    err := json.Unmarshal(jsonRequest, &request)
    if err != nil {
        return respondMalformed(nil)
    }
    attackErr := game.Attack(playerId, request.Attacker, request.AttackIndex, request.Target)
    if attackErr != nil {
        return respondBadCommand(attackErr.Error())
    }
    return respondSuccess(nil)
}

func endTurnCommand(jsonRequest []byte, playerId int, game *game_engine.Game) []byte {
    var request requests.MoveCommandRequest
    err := json.Unmarshal(jsonRequest, &request)
    if err != nil {
        return respondMalformed(nil)
    }
    endErr := game.EndTurn(playerId)
    if endErr != nil {
        return respondBadCommand(endErr.Error())
    }
    return respondSuccess(nil)
}

func exitCommand(jsonRequest []byte, playerId int, game *game_engine.Game) []byte {
    var request requests.ExitCommandRequest
    err := json.Unmarshal(jsonRequest, &request)
    if err != nil {
        return respondMalformed(nil)
    }
    return respondSuccess(nil)
}
