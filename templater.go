package main

import (
	"bytes"
	"dummy_client"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	err := os.Mkdir("templates/", 0777)
	if err != nil {
		fmt.Println("Couldn't make templates/ :", err)
	}
	conn, err := net.Dial("tcp", "localhost:5269")
	if err != nil {
		panic(err)
	}
	client := dummy_client.NewClient(conn)
	message, err := client.NewGame()
	if err != nil {
		panic(err)
	}
	writeJson(message)
	message, err = client.ViewWorld()
	if err != nil {
		panic(err)
	}
	writeJson(message)
	message, err = client.ViewPlayers()
	if err != nil {
		panic(err)
	}
	writeJson(message)
	message, err = client.ViewUnits()
	if err != nil {
		panic(err)
	}
	writeJson(message)
	message, err = client.ViewTerrain()
	if err != nil {
		panic(err)
	}
	writeJson(message)
	message, err = client.Move()
	if err != nil {
		panic(err)
	}
	writeJson(message)
	message, err = client.Attack()
	if err != nil {
		panic(err)
	}
	writeJson(message)
	message, err = client.Turn()
	if err != nil {
		panic(err)
	}
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
