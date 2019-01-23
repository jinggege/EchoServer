package main

import (
	"fmt"
	"strings"
)

func TelnetCmd(str string, exitChan chan int) bool {

	if strings.HasPrefix(str, "@close") {
		fmt.Println("=service will close")
		exitChan <- 0
		return false
	}

	return true
}
