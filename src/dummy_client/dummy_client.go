package dummy_client

import (
    "fmt"
    "net"
    "encoding/json"
    "requests"
    "github.com/DomoCo/connection"
)

type Client struct {
    conn connection.Connection
}

func NewClient(conn net.Conn) *Client {
    return &Client{conn: connection.NewSocketConn(conn)}
}

func (dummy Client) NewGame() (string, error) {
    return dummy.send("new", &requests.NewCommandRequest{Uids: []int{26, 13}})
}

func (dummy Client) View() (string, error) {
    return dummy.send("view:26", &requests.ViewCommandRequest{})
}

func (dummy Client) Move() (string, error) {
    return dummy.send("move:26", &requests.MoveCommandRequest{
        Move: []requests.LocationStruct{
            requests.LocationStruct{X: 0, Y: 0},
            requests.LocationStruct{X: 0, Y: 1}}})
}

func (dummy Client) Attack() (string, error) {
    return dummy.send("attack:26", &requests.AttackCommandRequest{
        Attacker: requests.LocationStruct{X: 0, Y: 0},
        AttackIndex: 0,
        Target: requests.LocationStruct{X: 0, Y: 1}})
}

func (dummy Client) Turn() (string, error) {
    return dummy.send("turn:26", &requests.EndTurnCommandRequest{PlayerId: 26})
}

func (dummy Client) Exit() (string, error) {
    return dummy.send("exit:26", &requests.ExitCommandRequest{Reason: "gameover"})
}

func (dummy Client) send(command string, jsonStringMessage interface{}) (string, error) {
    jsonMessage, err := json.Marshal(jsonStringMessage)
    if err != nil {
        return "", err
    }
    message := fmt.Sprintf("%s:%s", command, jsonMessage)
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
