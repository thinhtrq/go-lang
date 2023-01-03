package main

import "fmt"

type Person struct {
	name    string
	age     int
	address []Address
}

type Address struct {
	name   string
	number int
}

func main() {
	var person1 Person
	person1.age = 1
	person1.name = "TRAN QUOC THINH"
	person1.address = []Address{{"NHA1", 1}, {"NHA2", 2}}

	fmt.Println(person1)

	person2 := Person{
		name:    "THINH",
		age:     1,
		address: []Address{{"NHA1", 1}, {"NHA2", 2}},
	}

	fmt.Println(person2)
}
