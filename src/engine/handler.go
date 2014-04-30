package engine

import (
    "strconv"
    "bytes"
    "engine/game_engine"
)

type requestHandler struct {
    sessionGame *game_engine.Game
    gameRequest map[string]func([]byte, int, *game_engine.Game) []byte
}

func newRequestHandler() *requestHandler {
    return &requestHandler{
        sessionGame: nil,
        gameRequest: map[string]func([]byte, int, *game_engine.Game) []byte {
            "exit": exitRequest,
            "move": moveRequest,
            "turn": endTurnRequest,
            "attack": attackRequest,
            "view": viewRequest}}
}

func (handler *requestHandler) handleRequest(request []byte) []byte {
    command, requestJson := splitOnce(request)
    if handler.sessionGame == nil {
        if string(command) == "new" {
            response, game := newRequest(requestJson)
            handler.sessionGame = game
            return buildResponse(command, response)
        } else {
            return buildResponse(command, respondUnknownRequest("Need new game request"))
        }
    } else {
        fun, ok := handler.gameRequest[string(command)]; if ok {
            playerId, requestJsonNoPlayerId := splitOnce(requestJson)
            playerIdInt, err := strconv.Atoi(string(playerId))
            if err != nil {
                return buildResponse(command, respondMalformed("playerId not an int"))
            }
            response := fun(requestJsonNoPlayerId, playerIdInt, handler.sessionGame)
            return buildResponse(command, response)
        } else {
            return buildResponse(command, respondUnknownRequest("Unknown command"))
        }
    }
}

func buildResponse(command []byte, response []byte) []byte {
    return append(append(command, []byte(":")...), response...)
}

func splitOnce(input []byte) ([]byte, []byte) {
    pieces := bytes.SplitN(input, []byte(":"), 2)
    if len(pieces) == 1 {
        return pieces[0], []byte{}
    } else if len(pieces) == 2 {
        return pieces[0], pieces[1]
    } else {
        return []byte{}, []byte{}
    }
}
