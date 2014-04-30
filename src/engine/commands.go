package engine

import (
    "encoding/json"
    "api"
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

func respondBadRequest(payload interface{}) []byte {
    return generateResponse(payload, 3)
}

func respondUnknownRequest(payload interface{}) []byte {
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

func newRequest(jsonRequest []byte) ([]byte, *game_engine.Game) {
    var request api.NewRequest
    err := json.Unmarshal(jsonRequest, &request)
    if err != nil {
        return respondMalformed(nil), nil
    }
    game, err := game_engine.NewGame(request.Uids, request.Debug)
    if err != nil {
        return respondBadRequest(err.Error()), nil
    } else {
        return respondSuccess(nil), game
    }
}


func viewRequest(jsonRequest []byte, playerId int, game *game_engine.Game) []byte {
    var request api.ViewRequest
    err := json.Unmarshal(jsonRequest, &request)
    if err != nil {
        return respondMalformed(nil)
    }
    response, err := game.Serialize(playerId)
    if err != nil {
        return respondBadRequest(err.Error())
    }
    return respondSuccess(response)
}

func moveRequest(jsonRequest []byte, playerId int, game *game_engine.Game) []byte {
    var request api.MoveRequest
    err := json.Unmarshal(jsonRequest, &request)
    if err != nil {
        return respondMalformed(nil)
    }
    response, err := game.MoveUnit(playerId, request.Move)
    if err != nil {
        return respondBadRequest(err.Error())
    }
    return respondSuccess(response)
}

func attackRequest(jsonRequest []byte, playerId int, game *game_engine.Game) []byte {
    var request api.AttackRequest
    err := json.Unmarshal(jsonRequest, &request)
    if err != nil {
        return respondMalformed(nil)
    }
    response, err := game.Attack(playerId, request.Attacker, request.AttackIndex, request.Target)
    if err != nil {
        return respondBadRequest(err.Error())
    }
    return respondSuccess(response)
}

func endTurnRequest(jsonRequest []byte, playerId int, game *game_engine.Game) []byte {
    var request api.MoveRequest
    err := json.Unmarshal(jsonRequest, &request)
    if err != nil {
        return respondMalformed(nil)
    }
    response, err := game.EndTurn(playerId)
    if err != nil {
        return respondBadRequest(err.Error())
    }
    return respondSuccess(response)
}

func exitRequest(jsonRequest []byte, playerId int, game *game_engine.Game) []byte {
    var request api.ExitRequest
    err := json.Unmarshal(jsonRequest, &request)
    if err != nil {
        return respondMalformed(nil)
    }
    return respondSuccess(nil)
}
