package piscine

func StrRev(s string) string {
	slice := []rune(s) // turn string into a rune

	var empty []rune // create an empty rune

	for i := len(slice) - 1; i >= 0; i-- { // initialising from end of the string; the condition i >= 0 ; work backwards
		empty = append(empty, slice[i])
	}
	return string(empty) // print reverse
}

/*
package main

import (
	"fmt"
)

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func main() {
	fmt.Println(reverse("Hello World"))
}

*/
