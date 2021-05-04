package main

import (
	"fmt"
)

func main() {
	//fmt.Println("123")
	//fnWhile()
	//Foreach()
	//fnWhileUsingBreak()
	//fnPointer()
	//fnArray()
	fnIfElse()
}

func fnFor() {
	for index := 0; index <= 10; index++ {
		fmt.Print("Index %d\n", index)
	}
}

func fnWhile() {
	index := 0
	for index < 5 {
		index++
		fmt.Print("Index %d\n", index)
	}
}

func Foreach() {
	courses := []string{"Androind", "ios", "React"}

	for index, item := range courses {
		fmt.Printf("%d. %s\n", index+1, item)
	}

	for _, item := range courses {
		fmt.Printf("%s\n", item)
	}
}

func fnWhileUsingBreak() {
	index := 0
	for true {
		if index > 5 {
			break
		}

		fmt.Printf("WhileBreake-Index %d\n", index)
		index++
	}
}

func fnPointer() {
	msg := "sometime"
	var msgPionter *string = &msg

	fmt.Println(*msgPionter)

	fnChangeMessage(&msg, "new msg 1") //solution 1
	fmt.Println(msg)

	fnChangeMessage(msgPionter, "new msg 2") //solution 2 same solution 1 pass by & both fn
	fmt.Println(msg)

	fnChangeMessage(msgPionter, "new msg 3") //solution 3
	fmt.Println(*msgPionter)
}

func fnChangeMessage(aPointer *string, newMessege string) {
	*aPointer = "new msg"
}

func fnArray() {
	var arr []int = []int{1, 2, 3, 4}
	// var arr = []int{1,2,3,4}

	for _, item := range arr {
		fmt.Print(item, ", ")
	}
}

func fnIfElse() {
	somValues := 10
	if somValues == 10 {
		fmt.Println("==10")
	} else {
		fmt.Println("!=10")
	}

	if somValues > 10 || somValues < 2 {
		fmt.Printf("101")
	} else {
		fmt.Printf("202")
	}

	//short form
	if result := doSomething(); result == "OKs" {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}

func doSomething() string {
	return "OKs"
}
