package engine

import (
    "strconv"
    "bytes"
    "engine/game_engine"
)

type requestHandler struct {
    sessionGame *game_engine.Game
    gameCommand map[string]func([]byte, int, *game_engine.Game) []byte
}

func newRequestHandler() *requestHandler {
    return &requestHandler{
        sessionGame: nil,
        gameCommand: map[string]func([]byte, int, *game_engine.Game) []byte {
            "exit": exitCommand,
            "move": moveCommand,
            "turn": endTurnCommand,
            "attack": attackCommand,
            "view": viewCommand}}
}

func (handler *requestHandler) handleRequest(request []byte) []byte {
    command, requestJson := splitOnce(request)
    if handler.sessionGame == nil {
        if string(command) == "new" {
            response, game := newCommand(requestJson)
            handler.sessionGame = game
            return buildResponse(command, response)
        } else {
            return buildResponse(command, respondUnknownCommand("Need new game request"))
        }
    } else {
        fun, ok := handler.gameCommand[string(command)]; if ok {
            playerId, requestJsonNoPlayerId := splitOnce(requestJson)
            playerIdInt, err := strconv.Atoi(string(playerId))
            if err != nil {
                return buildResponse(command, respondMalformed("playerId not an int"))
            }
            response := fun(requestJsonNoPlayerId, playerIdInt, handler.sessionGame)
            return buildResponse(command, response)
        } else {
            return buildResponse(command, respondUnknownCommand("Unknown command"))
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
