package main

import "fmt"

var number int // use inside and outside function

func main() {
	var name string = "Tran quoc thinh"

	var name2 = "Tran quoc thinh 2"

	age := 20 // using inside function

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(age)

	var name3 string
	var age1 int
	var isBoy bool
	var point float64

	fmt.Println(name3)
	fmt.Println(age1)
	fmt.Println(isBoy)
	fmt.Println(point)

	name3 = "TRAN QUOC THINH 3"
	fmt.Println(name3)

	fmt.Println(number)

	number = 20
	fmt.Println(number)

	// declaration multiple variable
	var a, b, c, d int = 1, 2, 3, 4
	fmt.Println(a, b, c, d)

	// declaration block variable

	var (
		e int
		f string = "HUHU"
	)
	fmt.Println(e, f)

	/*


		A variable name must start with a letter or an underscore character (_)
		A variable name cannot start with a digit
		A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
		Variable names are case-sensitive (age, Age and AGE are three different variables)
		There is no limit on the length of the variable name
		A variable name cannot contain spaces
		The variable name cannot be any Go keywords
	*/
}
