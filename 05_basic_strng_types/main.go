package main

import (
	"fmt"
	"strings"
)

func main() {
	FirstName := "Aditya"
	LasName := "Ishan"
	FullName := FirstName + " " + LasName
	fmt.Println(FullName)

	fmt.Println(strings.ToUpper(FullName))
}
