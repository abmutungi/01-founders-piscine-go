package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		os.Args = os.Args[1:]
	} else {
		fmt.Println("File name missing")
	}

	if len(os.Args) == 1 {

		file, err := os.Open(os.Args[0])
		if err != nil {
			fmt.Println(err)
		}

		arr := make([]byte, 14)

		file.Read(arr)
		fmt.Println(string(arr))
		file.Close()
	} else {
		fmt.Println("Too many arguments")
	}
}
