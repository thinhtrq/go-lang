package main

import "fmt"

func main() {
	var arrNam = [3]string{"Tran", "Quoc", "Thinh"}
	fmt.Println(arrNam)

	arrAge := []int{1, 2, 3, 4}
	fmt.Println(arrAge)

	fmt.Println(len(arrAge))

	arrAge = append(arrAge, 5)
	fmt.Println(arrAge)
}
