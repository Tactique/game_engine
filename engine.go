package main

import (
	"engine"
	"fmt"
	"os"
	"strconv"
)

func main() {
	port := 5269
	if len(os.Args) == 2 {
		param_port, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		port = param_port
	}
	fmt.Printf("Starting Tactique on port %d!\n", port)
	engine.ListenForever(port)
}
