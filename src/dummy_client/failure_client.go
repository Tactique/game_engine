package dummy_client

import (
	"api"
	"fmt"
	"net"
)

func (dummy Client) BadNewGame() ([]byte, error) {
	return dummy.send(PanicSerialize("new", &api.NewRequest{Uids: []int{}}))
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
	fmt.Println(string(message))
}
