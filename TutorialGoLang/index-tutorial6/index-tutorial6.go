package main

import (
	"tutorial5/employee"
)

func main() {
	e := employee.Init(
		"Sam",
		"Adolf",
		30,
		20,
	)

	e.LeavesRemaining()

	e = employee.Init(
		"Am",
		"Adolf",
		30,
		20,
	)

	e.LeavesRemaining()
}
