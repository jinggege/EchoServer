package main

import (
	"os"
)

func start() {
	exitChan := make(chan int)

	TS := &TelnetService{}

	go TS.Run("127.0.0.1:3001", exitChan)

	code := <-exitChan

	os.Exit(code)
}

func main() {
	start()
	for {
	}
}
