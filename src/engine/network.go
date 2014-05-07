package engine

import (
	"fmt"
	"net"
	"github.com/Tactique/golib/connection"
	"github.com/Tactique/golib/logger"
)

func ListenForever(port int) error {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Warnf("Error accepting connection (%s)", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(netConn net.Conn) error {
	conn := connection.NewSocketConn(netConn)
	handler := NewRequestHandler()
	for {
		logger.Infof("Handling a connection")
		request, err := conn.Read()
		if err != nil {
			logger.Warnf("Error reading request (%s)", err.Error())
			conn.Close()
			return err
		}
		logger.Debugf("Got request %s", string(request))
		response := handler.HandleRequest(request)
		logger.Debugf("Sent response %s", string(response))
		err = conn.Write(response)
		if err != nil {
			logger.Warnf("Error writing response (%s)", err.Error())
			conn.Close()
			return err
		}
	}
}
