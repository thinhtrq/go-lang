package main

import (
	"fmt"
)

func sum(num1 int, num2 float64) float64 {
	return num2 + float64(num1)
}

func devices(num1 int, num2 int) (int, int, error) {

	return 2, num1 / num2, nil
}

func main() {
	fmt.Println(sum(1, 2.3))
	b, c, err := devices(1, 0)

	fmt.Printf("%v--hh--%d---%v\n", c, b, err)
}
