package engine

import (
	"dummy_client"
	"encoding/json"
	"io/ioutil"
	"path"
)

func generateTemplates(filepath string) {
	handler := newRequestHandler()

	message := handler.handleRequest(dummy_client.BuiltNewRequest)
	writeJson(filepath, message)
	message = handler.handleRequest(dummy_client.BuiltViewWorldRequest)
	writeJson(filepath, message)
	message = handler.handleRequest(dummy_client.BuiltViewTerrainRequest)
	writeJson(filepath, message)
	message = handler.handleRequest(dummy_client.BuiltViewUnitsRequest)
	writeJson(filepath, message)
	message = handler.handleRequest(dummy_client.BuiltViewPlayersRequest)
	writeJson(filepath, message)
	message = handler.handleRequest(dummy_client.BuiltMoveRequest)
	writeJson(filepath, message)
	message = handler.handleRequest(dummy_client.BuiltAttackRequest)
	writeJson(filepath, message)
	message = handler.handleRequest(dummy_client.BuiltEndTurnRequest)
	writeJson(filepath, message)
}

func writeJson(filepath string, message []byte) {
	name, contents := splitOnce(message)
	writableContents := transformResponse(name, contents)
	err := ioutil.WriteFile(path.Join(filepath, string(name)+".json"), writableContents, 0644)
	if err != nil {
		panic(err)
	}
}

func transformResponse(name []byte, message []byte) []byte {
	response := make(map[string]interface{}, 0)
	err := json.Unmarshal(message, &response)
	if err != nil {
		panic(err)
	}
	writableResponse, err := json.MarshalIndent(map[string]interface{}{
		string(name): response["payload"]}, "", "    ")
	if err != nil {
		panic(err)
	}
	return writableResponse
}
