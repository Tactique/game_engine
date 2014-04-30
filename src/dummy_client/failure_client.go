package dummy_client

import (
    "fmt"
    "net"
    "api"
)

func (dummy Client) BadNewGame() (string, error) {
    return dummy.send("new", &api.NewRequest{Uids: []int{}})
}

func BadTestRun() {
    conn, err := net.Dial("tcp", "localhost:5269")
    if err != nil {
        panic(err)
    }
    client := NewClient(conn)
    message, err := client.BadNewGame()
    if err != nil {
        panic(err)
    }
    fmt.Println(message)
}
