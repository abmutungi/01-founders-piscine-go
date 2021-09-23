package main

import (
	"os"
)

// Write a program that prints the arguments receivedin the command line

func main() {
	arguments := os.Args
	count := 0
	// Function body is called upon 'arguments' above, has to be in main file
	// Convention go run 'x.go' argument , new argument indexed after any space
	// Declare argument in body
	// Loop/iterate to convert string to rune to ascertain values

	for i := 1; i < len(arguments); i++ { // len(args) gives number of arguments
		for _, elements := range arguments {
			count = count + len(arguments)
		}
		return
	}
}

/*for _, p := range f.Params {
+		n += len(p.SyscallArgList())
+	}
+	return n
+}
*/
