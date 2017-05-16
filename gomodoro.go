package main

import (
	"fmt"
	s "strings"
	"time"
)

var p = fmt.Print

func main() {
	fmt.Println("25 min")
	cnt := 0
	for {
		p(s.Repeat("\b", 25))
		p(s.Repeat("=", cnt))
		p(">")
		if 24-cnt > 1 {
			p(s.Repeat("-", 24-cnt))
		}

		time.Sleep(1 * time.Second)
		cnt++
		if 1*cnt > 24 {
			fmt.Println("\a")
			break
		}
	}
}

func message() string {
	return "hello"
}
