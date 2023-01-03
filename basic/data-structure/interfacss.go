package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type Student interface {
	goToSchool() string
}

type Nam struct {
	schoolName string
	address    string
}

type Nu struct {
	schoolName string
	address    string
}

func (nam Nam) goToSchool() string {
	return nam.schoolName
}

func (nu Nu) goToSchool() string {
	return nu.address
}

func goToSchool(student Student) string {
	return student.goToSchool()
}

func Sqrt(value float64) (result float64, err error) {
	if value < 0 {
		return 0, errors.New("ddddd")
	}
	return math.Sqrt(value), nil
}

func main() {
	nam := Nam{
		schoolName: "FFFF",
		address:    "DDDDD",
	}

	fmt.Println(goToSchool(nam))

	_, err := Sqrt(-1)
	fmt.Println(strings.Split(err.Error(), ","))
}
