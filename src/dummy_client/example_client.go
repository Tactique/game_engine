package dummy_client

import (
	"fmt"
	"net"
)

func ExampleTestRun() {
	conn, err := net.Dial("tcp", "localhost:5269")
	if err != nil {
		panic(err)
	}
	proxy := NewClient(conn)
	message, err := proxy.NewGame()
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
	message, err = proxy.ViewWorld()
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
	message, err = proxy.Move()
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
	message, err = proxy.ViewWorld()
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
	message, err = proxy.Attack()
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
	message, err = proxy.ViewWorld()
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
	message, err = proxy.Turn()
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
	message, err = proxy.ViewWorld()
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
