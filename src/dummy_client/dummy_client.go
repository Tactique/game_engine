package dummy_client

import (
	"api"
	"encoding/json"
	"fmt"
	"github.com/DomoCo/connection"
	"net"
)

type Client struct {
	conn connection.Connection
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn: connection.NewSocketConn(conn)}
}

func (dummy Client) NewGame() (string, error) {
	return dummy.send("new", &api.NewRequest{Uids: []int{26, 13}})
}

func (dummy Client) ViewWorld() (string, error) {
	return dummy.send("viewWorld:26", &api.ViewWorldRequest{})
}

func (dummy Client) ViewTerrain() (string, error) {
	return dummy.send("viewTerrain:26", &api.ViewTerrainRequest{})
}

func (dummy Client) ViewUnits() (string, error) {
	return dummy.send("viewUnits:26", &api.ViewUnitsRequest{})
}

func (dummy Client) ViewPlayers() (string, error) {
	return dummy.send("viewPlayers:26", &api.ViewPlayersRequest{})
}

func (dummy Client) Move() (string, error) {
	return dummy.send("move:26", &api.MoveRequest{
		Move: []api.LocationStruct{
			api.LocationStruct{X: 0, Y: 0},
			api.LocationStruct{X: 0, Y: 1}}})
}

func (dummy Client) Attack() (string, error) {
	return dummy.send("attack:26", &api.AttackRequest{
		Attacker:    api.LocationStruct{X: 0, Y: 1},
		AttackIndex: 0,
		Target:      api.LocationStruct{X: 0, Y: 3}})
}

func (dummy Client) Turn() (string, error) {
	return dummy.send("turn:26", &api.EndTurnRequest{})
}

func (dummy Client) Exit() (string, error) {
	return dummy.send("exit:26", &api.ExitRequest{Reason: "gameover"})
}

func (dummy Client) send(command string, jsonStringMessage interface{}) (string, error) {
	fmt.Printf("Sending %v\n", jsonStringMessage)
	jsonMessage, err := json.Marshal(jsonStringMessage)
	if err != nil {
		return "", err
	}
	message := fmt.Sprintf("%s:%s", command, jsonMessage)
	fmt.Printf("Sending %s\n", message)
	err = dummy.conn.Write([]byte(message))
	if err != nil {
		return "", err
	}
	response, err := dummy.conn.Read()
	if err != nil {
		return "", err
	}
	return string(response), nil
}
