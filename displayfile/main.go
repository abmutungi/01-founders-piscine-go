package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		os.Args = os.Args[1:]
	} else {
		fmt.Println("File name missing")
		return
	}

	if len(os.Args) == 1 {

		file, data := os.Open(os.Args[0])
		if data != nil {
			fmt.Println(data)
		}

		arr := make([]byte, 14)

		file.Read(arr)
		fmt.Println(string(arr))
		file.Close()
	} else {
		fmt.Println("Too many arguments")
	}
}
