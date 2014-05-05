package dummy_client

import (
	"api"
	"encoding/json"
	"fmt"
	"github.com/Tactique/golib/connection"
	"net"
)

var BuiltNewRequest []byte = PanicSerialize("new", &api.NewRequest{Uids: []int{26, 13}})

var BuiltViewWorldRequest []byte = PanicSerialize("viewWorld:26", &api.ViewWorldRequest{})

var BuiltViewTerrainRequest []byte = PanicSerialize("viewTerrain:26", &api.ViewTerrainRequest{})

var BuiltViewUnitsRequest []byte = PanicSerialize("viewUnits:26", &api.ViewUnitsRequest{})

var BuiltViewPlayersRequest []byte = PanicSerialize("viewPlayers:26", &api.ViewPlayersRequest{})

var BuiltMoveRequest []byte = PanicSerialize("move:26", &api.MoveRequest{
		Move: []api.LocationStruct{
			api.LocationStruct{X: 0, Y: 0},
			api.LocationStruct{X: 0, Y: 1}}})

var BuiltAttackRequest []byte = PanicSerialize("attack:26", &api.AttackRequest{
		Attacker:    api.LocationStruct{X: 0, Y: 1},
		AttackIndex: 0,
		Target:      api.LocationStruct{X: 0, Y: 3}})

var BuiltEndTurnRequest []byte = PanicSerialize("turn:26", &api.EndTurnRequest{})

var BuiltExitRequest []byte = PanicSerialize("exit:26", &api.ExitRequest{Reason: "gameover"})

type Client struct {
	conn connection.Connection
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn: connection.NewSocketConn(conn)}
}

func (dummy Client) NewGame() ([]byte, error) {
	return dummy.send(BuiltNewRequest)
}

func (dummy Client) ViewWorld() ([]byte, error) {
	return dummy.send(BuiltViewWorldRequest)
}

func (dummy Client) ViewTerrain() ([]byte, error) {
	return dummy.send(BuiltViewTerrainRequest)
}

func (dummy Client) ViewUnits() ([]byte, error) {
	return dummy.send(BuiltViewUnitsRequest)
}

func (dummy Client) ViewPlayers() ([]byte, error) {
	return dummy.send(BuiltViewPlayersRequest)
}

func (dummy Client) Move() ([]byte, error) {
	return dummy.send(BuiltMoveRequest)
}

func (dummy Client) Attack() ([]byte, error) {
	return dummy.send(BuiltAttackRequest)
}

func (dummy Client) Turn() ([]byte, error) {
	return dummy.send(BuiltEndTurnRequest)
}

func (dummy Client) Exit() ([]byte, error) {
	return dummy.send(BuiltExitRequest)
}

func buildRequest(command string, response []byte) []byte {
	return append(append([]byte(command), byte(':')), response...)
}

func PanicSerialize(command string, jsonStringMessage interface{}) []byte {
	response, err := Serialize(command , jsonStringMessage)
	if err != nil {
		panic(err)
	}
	return response
}


func Serialize(command string, jsonStringMessage interface{}) ([]byte, error) {
	jsonMessage, err := json.Marshal(jsonStringMessage)
	if err != nil {
		return []byte{}, err
	}
	return buildRequest(command, jsonMessage), nil
}


func (dummy Client) send(message []byte) ([]byte, error) {
	fmt.Printf("Sending %s\n", message)
	err := dummy.conn.Write([]byte(message))
	if err != nil {
		return []byte{}, err
	}
	response, err := dummy.conn.Read()
	if err != nil {
		return []byte{}, err
	}
	return response, nil
}
