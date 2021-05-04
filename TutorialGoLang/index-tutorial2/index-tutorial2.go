package main

import (
	"fmt"
)

func main() {
	//var number []int
	//numbers = append(number, 1)
	var numbers1 = make([]int, 0, 5) //Slice
	numbers1 = append(numbers1, 1)
	numbers1 = append(numbers1, 2)
	showSlice(numbers1)

	var number2 []int
	number2 = append(number2, 1)
	number2 = append(number2, 2)
	showSlice(number2)

	var numbers = []int{1, 2, 3}
	//leading remove
	numbers = numbers[1:len(numbers)]
	//trailing remove
	numbers = numbers[0 : len(numbers)-1]
	showSlice(numbers)

	//remove specific index
	//numbers = removeIndex(numbers, 1)
	//showSlice(numbers)

	fnMap()
	fnMapDynamic()
}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...) // 0:index , index+1:
} // spreding all of information in array

func fnMap() {
	var number4 = map[string]int{"one": 1, "two": 2, "tree": 3, "four": 4}
	fmt.Println("", number4)
}

func fnMapDynamic() {
	var colors = make(map[string]string)
	colors["red"] = "#f00"
	colors["green"] = "#f00"
	colors["blue"] = "#f00"
	fmt.Println(colors)

	var courses = make(map[string]map[string]int)
	//colorses["Android"] = map[string]int{"price": 20} or
	courses["Android"] = make(map[string]int)

	courses["Android"]["code"] = 1234
	fmt.Println(courses)
}
