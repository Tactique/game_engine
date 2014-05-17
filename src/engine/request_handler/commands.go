package request_handler

import (
	"api"
	"encoding/json"
	"github.com/Tactique/golib/logger"
)

func respondSuccess(payload interface{}) []byte {
	return generateResponse(payload, api.STATUS_OK)
}

func respondFailure(payload interface{}) []byte {
	return generateResponse(payload, api.STATUS_FAILURE)
}

func respondMalformed(payload interface{}) []byte {
	return generateResponse(payload, api.STATUS_MALFORMED_RESPONSE)
}

func respondBadRequest(payload interface{}) []byte {
	return generateResponse(payload, api.STATUS_BAD_REQUEST)
}

func respondUnknownRequest(payload interface{}) []byte {
	return generateResponse(payload, api.STATUS_UNKNOWN_REQUEST)
}

func generateResponse(payload interface{}, status int) []byte {
	response, err := json.Marshal(api.ResponseType{Status: status, Payload: payload})
	if err != nil {
		logger.Warnf("Could not generate response: %s", err.Error())
		backupResponse, err := json.Marshal(api.ResponseType{
			Status:  api.STATUS_UNSERIALIZEABLE_RESPONSE,
			Payload: "problem"})
		if err != nil {
			return []byte("Really bad")
		}
		logger.Infof("Generating response with status %d", status)
		logger.Debugf("Full message %s", string(response))
		return backupResponse
	}
	logger.Infof("Generating response with status %d", status)
	logger.Debugf("Full message %s", string(response))
	return response
}

func newRequest(jsonRequest []byte) ([]byte, *GameWrapper) {
	var request api.NewRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil), nil
	}
	game, err := NewGameWrapper(request)
	if err != nil {
		return respondBadRequest(err.Error()), nil
	} else {
		return respondSuccess(nil), game
	}
}

func viewWorldRequest(jsonRequest []byte, playerId int, game *GameWrapper) []byte {
	var request api.ViewWorldRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil)
	}
	response, err := game.ViewWorld(playerId, request)
	if err != nil {
		return respondBadRequest(err.Error())
	}
	return respondSuccess(response)
}

func viewPlayersRequest(jsonRequest []byte, playerId int, game *GameWrapper) []byte {
	var request api.ViewPlayersRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil)
	}
	response, err := game.ViewPlayers(playerId, request)
	if err != nil {
		return respondBadRequest(err.Error())
	}
	return respondSuccess(response)
}

func viewUnitsRequest(jsonRequest []byte, playerId int, game *GameWrapper) []byte {
	var request api.ViewUnitsRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil)
	}
	response, err := game.ViewUnits(playerId, request)
	if err != nil {
		return respondBadRequest(err.Error())
	}
	return respondSuccess(response)
}

func viewTerrainRequest(jsonRequest []byte, playerId int, game *GameWrapper) []byte {
	var request api.ViewTerrainRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil)
	}
	response, err := game.ViewTerrain(playerId, request)
	if err != nil {
		return respondBadRequest(err.Error())
	}
	return respondSuccess(response)
}

func moveRequest(jsonRequest []byte, playerId int, game *GameWrapper) []byte {
	var request api.MoveRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil)
	}
	response, err := game.MoveUnit(playerId, request)
	if err != nil {
		return respondBadRequest(err.Error())
	}
	return respondSuccess(response)
}

func attackRequest(jsonRequest []byte, playerId int, game *GameWrapper) []byte {
	var request api.AttackRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil)
	}
	response, err := game.Attack(playerId, request)
	if err != nil {
		return respondBadRequest(err.Error())
	}
	return respondSuccess(response)
}

func endTurnRequest(jsonRequest []byte, playerId int, game *GameWrapper) []byte {
	var request api.EndTurnRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil)
	}
	response, err := game.EndTurn(playerId, request)
	if err != nil {
		return respondBadRequest(err.Error())
	}
	return respondSuccess(response)
}

func exitRequest(jsonRequest []byte, playerId int, game *GameWrapper) []byte {
	var request api.ExitRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil)
	}
	return respondSuccess(nil)
}
