package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// range returns two values in each iteration
// index, and the value at index
// we don't need index in this program, so we assign it to a blank identifier
// and the value at index to arg

// The following statements are all the same
// s := ""						can only be used within a function (not for global variables)
// var s string				use when you need to initialize the variable to 0 or ""
// var s = ""					use when you need to declare, and initialize multiple variables
// var s string = "" 	use when you need to explicityly specify the type of a variable when declaring, and initializing multiple variables
