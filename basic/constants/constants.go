package main

import (
	"fmt"
)

const CONTRY_SIDE = "VIET NAME"

func main() {

	fmt.Println(CONTRY_SIDE)

	// multiple constants declaration

	const (
		A = 1
		B = "2"
		C = true
	)

	fmt.Println(A, B, C)
}
