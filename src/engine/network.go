package engine

import (
	"fmt"
	"github.com/DomoCo/connection"
	"net"
)

func ListenForever(port int) error {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(netConn net.Conn) error {
	conn := connection.NewSocketConn(netConn)
	handler := newRequestHandler()
	for {
		fmt.Println("Handling a connection")
		request, err := conn.Read()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			return err
		}
		fmt.Println("Got request", string(request))
		response := handler.handleRequest(request)
		fmt.Println("Sent response", string(response))
		err = conn.Write(response)
		if err != nil {
			fmt.Println(err)
			conn.Close()
			return err
		}
	}
}
