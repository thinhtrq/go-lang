package main

import "fmt"

func main() {
	var arrAge = []int{0, 0, 0, 0}

	fmt.Println(arrAge)

	arrAge = append(arrAge, 1, 2, 3, 4)
	fmt.Println(arrAge)
	fmt.Printf("%v----%T", arrAge, arrAge)

	var myArr = []int{1, 2, 3, 4, 5, 6}

	mySlices := myArr[0 : len(myArr)-1]

	fmt.Println(mySlices)
}
