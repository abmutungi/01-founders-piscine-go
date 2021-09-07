package main

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb >= 0 {
		z01.Println("F")
	} else {
		z01.Println("T")
	}
}
