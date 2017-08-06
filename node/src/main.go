package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, let's download ;-)")

	for {
		time.Sleep(5 * time.Minute)
		fmt.Println("tita")
	}
}
