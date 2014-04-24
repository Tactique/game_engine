package main

import (
    "net"
    "fmt"
    "dummy_proxy"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:5269")
    if err != nil {
        panic(err)
    }
    proxy := dummy_proxy.NewDummyProxy(conn)
    message, err := proxy.BadNewGame()
    if err != nil {
        panic(err)
    }
    fmt.Println(message)
}
