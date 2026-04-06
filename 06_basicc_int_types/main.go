package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 20
	y++
	y++
	z := x + y ^ 2 // what is y^2 - is it 20^2 = 400 or is it 20 XOR 2 = 22?
	// in go, ^ is the bitwise XOR operator, so y^2 is 20 XOR 2 which equals 22
	fmt.Println(x, y, z)

	//In Go's operator precedence, addition (+) has higher precedence than bitwise XOR (^).
	// The expression is evaluated as: (x + y) ^ 2.
	// 10 + 22 = 32.
	// 32 ^ 2 (Bitwise XOR):
	// 32 in binary: 100000
	// 2 in binary:  000010
	// Result:       100010 (which is 34 in decimal).
	flt := 3.14
	fmt.Println(flt)
}

// other forms of operators are:
// x += 10 // x = x + 10
// x -= 10 // x = x - 10
// AND operator: &&
// OR operator: ||
// NOT operator: !
