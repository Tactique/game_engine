package engine

import (
    "strconv"
    "strings"
    "engine/game_engine"
)

type requestHandler struct {
    sessionGame *game_engine.Game
    newGameCommand map[string]func(string) (string, *game_engine.Game)
    gameCommand map[string]func(string, int, *game_engine.Game) string
}

func newRequestHandler() *requestHandler {
    return &requestHandler{
        sessionGame: nil,
        newGameCommand: map[string]func(string) (string, *game_engine.Game){
            "new": newCommand},
        gameCommand: map[string]func(string, int, *game_engine.Game) string {
            "exit": exitCommand,
            "move": moveCommand,
            "turn": endTurnCommand,
            "attack": attackCommand,
            "view": viewCommand}}
}

func (handler *requestHandler) handleRequest(request string) string {
    command, requestJson := splitOnce(request)
    if handler.sessionGame == nil {
        fun, ok := handler.newGameCommand[command]; if ok {
            response, game := fun(requestJson)
            handler.sessionGame = game
            return command + ":" + response
        } else {
            return command + ":" + respondUnknownCommand("Need new game request")
        }
    } else {
        fun, ok := handler.gameCommand[command]; if ok {
            playerId, requestJsonNoPlayerId := splitOnce(requestJson)
            playerIdInt, err := strconv.Atoi(playerId)
            if err != nil {
                return respondMalformed("playerId not an int")
            }
            response := fun(requestJsonNoPlayerId, playerIdInt, handler.sessionGame)
            return command + ":" + response
        } else {
            return command + ":" + respondUnknownCommand("Unknown command")
        }
    }
}

func splitOnce(inputString string) (string, string) {
    pieces := strings.SplitN(inputString, ":", 2)
    if len(pieces) == 1 {
        return pieces[0], ""
    } else if len(pieces) == 2 {
        return pieces[0], pieces[1]
    } else {
        return "", ""
    }
}
