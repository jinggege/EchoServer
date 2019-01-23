package main

import (
	"fmt"
	"net"
)

type TelnetService struct {
}

func (ts *TelnetService) Run(address string, exitChan chan int) {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}

	fmt.Println("==listen on :", address)

	defer listen.Close()

	for {

		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("==error accept ", err.Error())
			continue
		}

		go ParseSession(conn, exitChan)
	}

}
