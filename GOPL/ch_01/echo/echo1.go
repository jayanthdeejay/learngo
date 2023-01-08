// Echo1 prints it's command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// for is the only type of loop in Go
// for initialization; condition; post {
//   zero or more statements
// }

// initialization at the start of the loop
// condition before every loop
// post after every loop
// ++/-- cannot be prefixes like ++i (no no)

// a while loop
// for condition {
// ...
// }

// infinite loop
// for {
// ...
// }
