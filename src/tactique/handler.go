package tactique

import (
    "strconv"
    "strings"
    "tactique/game_engine"
)

type requestHandler struct {
    sessionGame *game_engine.Game
    newGameCommand map[string]func(string) (string, *game_engine.Game, error)
    gameCommand map[string]func(string, int, *game_engine.Game) (string, error)
}

func newRequestHandler() *requestHandler {
    return &requestHandler{
        sessionGame: nil,
        newGameCommand: map[string]func(string) (string, *game_engine.Game, error){
            "new": newCommand},
        gameCommand: map[string]func(string, int, *game_engine.Game) (string, error){
            "exit": exitCommand,
            "move": moveCommand,
            "turn": endTurnCommand,
            "view": viewCommand}}
}

func (handler *requestHandler) handleRequest(request string) (string, error) {
    command, requestJson := splitOnce(request)
    if handler.sessionGame == nil {
        fun, ok := handler.newGameCommand[command]; if ok {
            response, game, err := fun(requestJson)
            if err != nil {
                return response, err
            } else {
                handler.sessionGame = game
                return response, err
            }
        } else {
            return "Need new game request", nil
        }
    } else {
        fun, ok := handler.gameCommand[command]; if ok {
            playerId, requestJsonNoPlayerId := splitOnce(requestJson)
            playerIdInt, err := strconv.Atoi(playerId)
            if err != nil {
                return "", err
            }
            return fun(requestJsonNoPlayerId, playerIdInt, handler.sessionGame)
        } else {
            return "Unknown command", nil
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
