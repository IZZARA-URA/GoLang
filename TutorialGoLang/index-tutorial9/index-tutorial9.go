// Chanel

package main

import (
	"fmt"
	"time"
)

func routine(c chan int, payload int) {
	c <- payload
	fmt.Println("step1")
	c <- payload
	fmt.Println("step2")
}

func main() {
	ch := make(chan int, 1)
	go routine(ch, 1)

	fmt.Println(<-ch) // read ch
	time.Sleep(1 * time.Second)

}
