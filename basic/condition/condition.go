package main

import "fmt"

func main() {

	if 10 > 5 {
		fmt.Println("Number 10 greater than 5 ")

	}

	x := 2
	y := 5
	if x < y {
		fmt.Println("Number 2 less than 5 ")
	} else {
		//
	}

	switch x {
	case 2, 5, 6:
		fmt.Printf("Number %d \n", x)

	default:
		fmt.Printf("Number %d \n", x)

	}

	var age = []int{1, 2, 3, 4, 5, 5, 6, 7}

	for i := 0; i < len(age); i++ {
		fmt.Println(age[i])
	}

	for i := 0; i < len(age); i++ {
		defer fmt.Println(age[i])
	}

	for i := len(age) - 1; i >= 0; i-- {
		fmt.Println(age[i])
	}

	for index, x := range age {
		fmt.Printf("index %d----%d \n", index, x)
	}
}
