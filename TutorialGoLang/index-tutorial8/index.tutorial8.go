// Sequential vs Concurrent vs Parallal

package main

import (
	"fmt"
	"time"
)

func run1() {
	fmt.Println("Run something")
}

func run2() {
	fmt.Println("Run something")
}

func main() {
	go run1()
	go run2()

	time.Sleep(5 * time.Second)
}
