package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func ParseSession(conn net.Conn, exitChan chan int) {
	fmt.Println("=session start")

	reader := bufio.NewReader(conn)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("=session error ")
			conn.Close()
			break
		}
		fmt.Println("[msg from client]:", str)

		if !TelnetCmd(str, exitChan) {
			conn.Write([]byte("--conn will closed--"))
			conn.Close()
			break
		}
		str = "[replay]:" + strings.TrimSpace(str)
		conn.Write([]byte(str + "\r\n"))
	}

}
