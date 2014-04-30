package engine

import (
    "net"
    "fmt"
    "github.com/DomoCo/connection"
)

type proxyConnection struct {
    conn connection.Connection
    handler *requestHandler
}

func newProxyConnection (conn net.Conn) *proxyConnection {
    return &proxyConnection{conn: connection.NewSocketConn(conn), handler: newRequestHandler()}
}

func (tc *proxyConnection) Read() (string, error) {
    byteslice, err := tc.conn.Read()
    return string(byteslice), err
}

func (tc *proxyConnection) Write(response string) error {
    return tc.conn.Write([]byte(response))
}

func (tc *proxyConnection) Close() error {
    return tc.conn.Close()
}

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
    conn := newProxyConnection(netConn)
    for {
        fmt.Println("Handling a connection")
        request, err := conn.Read()
        if err != nil {
            fmt.Println(err)
            conn.Close()
            return err
        }
        fmt.Println("Got request", request)
        response := conn.handler.handleRequest(request)
        /*
        if err != nil {
            fmt.Println(err)
            conn.Close()
            return err
        }
        */
        fmt.Println("Sent response", response)
        err = conn.Write(response)
        if err != nil {
            fmt.Println(err)
            conn.Close()
            return err
        }
    }
}
