package dummy_client

import (
    "fmt"
    "net"
    "encoding/json"
    "api"
    "github.com/DomoCo/connection"
)

type Client struct {
    conn connection.Connection
}

func NewClient(conn net.Conn) *Client {
    return &Client{conn: connection.NewSocketConn(conn)}
}

func (dummy Client) NewGame() (string, error) {
    return dummy.send("new", &api.NewCommandRequest{Uids: []int{26, 13}})
}

func (dummy Client) View() (string, error) {
    return dummy.send("view:26", &api.ViewCommandRequest{})
}

func (dummy Client) Move() (string, error) {
    return dummy.send("move:26", &api.MoveCommandRequest{
        Move: []api.LocationStruct{
            api.LocationStruct{X: 0, Y: 0},
            api.LocationStruct{X: 0, Y: 1}}})
}

func (dummy Client) Attack() (string, error) {
    return dummy.send("attack:26", &api.AttackCommandRequest{
        Attacker: api.LocationStruct{X: 0, Y: 1},
        AttackIndex: 0,
        Target: api.LocationStruct{X: 0, Y: 3}})
}

func (dummy Client) Turn() (string, error) {
    return dummy.send("turn:26", &api.EndTurnCommandRequest{PlayerId: 26})
}

func (dummy Client) Exit() (string, error) {
    return dummy.send("exit:26", &api.ExitCommandRequest{Reason: "gameover"})
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
