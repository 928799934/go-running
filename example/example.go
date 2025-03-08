package main

import (
	"fmt"

	gorunning "github.com/928799934/go-running"
)

func main() {
	fmt.Println("start")
	gorunning.Loop(func() {
		fmt.Println("exit")
	}, func() {}, nil)
}
