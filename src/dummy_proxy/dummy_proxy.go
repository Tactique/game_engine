package dummy_proxy

import (
    "fmt"
    "net"
    "encoding/json"
    "requests"
    "github.com/DomoCo/connection"
)

type DummyProxy struct {
    conn connection.Connection
}

func NewDummyProxy(conn net.Conn) *DummyProxy {
    return &DummyProxy{conn: connection.NewSocketConn(conn)}
}

func (dummy DummyProxy) BadNewGame() (string, error) {
    return dummy.send("new", &requests.NewCommandRequest{Uids: []int{26}})
}

func (dummy DummyProxy) NewGame() (string, error) {
    return dummy.send("new", &requests.NewCommandRequest{Uids: []int{26, 13}})
}

func (dummy DummyProxy) View() (string, error) {
    return dummy.send("view", &requests.ViewCommandRequest{Uid: 26})
}

func (dummy DummyProxy) Move() (string, error) {
    return dummy.send("move", &requests.MoveCommandRequest{
        Move: []requests.LocationStruct{
            requests.LocationStruct{X: 0, Y: 0},
            requests.LocationStruct{X: 0, Y: 1}}})
}

func (dummy DummyProxy) Turn() (string, error) {
    return dummy.send("turn", &requests.EndTurnCommandRequest{PlayerId: 26})
}

func (dummy DummyProxy) Exit() (string, error) {
    return dummy.send("exit", &requests.ExitCommandRequest{Reason: "gameover"})
}

func (dummy DummyProxy) send(command string, jsonStringMessage interface{}) (string, error) {
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

func ExampleTestRun() {
    conn, err := net.Dial("tcp", "localhost:5269")
    if err != nil {
        panic(err)
    }
    proxy := NewDummyProxy(conn)
    message, err := proxy.NewGame()
    if err != nil {
        panic(err)
    }
    fmt.Println(message)
    message, err = proxy.View()
    if err != nil {
        panic(err)
    }
    fmt.Println(message)
    message, err = proxy.Move()
    if err != nil {
        panic(err)
    }
    fmt.Println(message)
    message, err = proxy.View()
    if err != nil {
        panic(err)
    }
    fmt.Println(message)
    message, err = proxy.Turn()
    if err != nil {
        panic(err)
    }
    fmt.Println(message)
    message, err = proxy.View()
    if err != nil {
        panic(err)
    }
    fmt.Println(message)
    message, err = proxy.Exit()
    if err != nil {
        panic(err)
    }
    fmt.Println(message)
}
