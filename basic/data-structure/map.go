package main

import "fmt"

func main() {
	var mapValue = map[string]string{}
	mapValue["ddd"] = "fff"

	fmt.Println(mapValue)

	var map1 = make(map[string]int)
	fmt.Println(map1)
}
