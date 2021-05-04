package main

import (
	"tutorial5/employee"
)

func main() {
	e := employee.New(
		"Sam",
		"Adolf",
		30,
		20,
	)

	e.LeavesRemaining()
}
