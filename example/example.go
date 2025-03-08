package main

import (
	"fmt"

	"github.com/928799934/go-running"
)

func main() {
	fmt.Println("start")
	running.Loop(func() {
		fmt.Println("exit")
	}, func() {}, nil)
}
