package main

import (
	"fmt"
)

func BasicAtoi(s string) int {
	var result int = 0

	for _, var num  = range s {

		var conv = int(num) - 48
		result = (result * 10) + conv
	}

	return result
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}

