package templater

import (
	"dummy_client"
	"engine"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func GenerateTemplates() {
	err := os.Mkdir("templates/", 0777)
	if err != nil {
		fmt.Println("Couldn't make templates/ :", err)
	}

	handler := engine.NewRequestHandler()

	message := handler.HandleRequest(dummy_client.BuiltNewRequest)
	writeJson(message)
	message = handler.HandleRequest(dummy_client.BuiltViewWorldRequest)
	writeJson(message)
	message = handler.HandleRequest(dummy_client.BuiltViewTerrainRequest)
	writeJson(message)
	message = handler.HandleRequest(dummy_client.BuiltViewUnitsRequest)
	writeJson(message)
	message = handler.HandleRequest(dummy_client.BuiltViewPlayersRequest)
	writeJson(message)
	message = handler.HandleRequest(dummy_client.BuiltMoveRequest)
	writeJson(message)
	message = handler.HandleRequest(dummy_client.BuiltAttackRequest)
	writeJson(message)
	message = handler.HandleRequest(dummy_client.BuiltEndTurnRequest)
	writeJson(message)
}

func writeJson(message []byte) {
	name, contents := splitOnce(message)
	writableContents := transformResponse(name, contents)
	err := ioutil.WriteFile("templates/"+string(name)+".json", writableContents, 0644)
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
