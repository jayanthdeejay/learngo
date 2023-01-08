// Echo to print index, and value of arguments, one per line
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Println(index, ":", arg)
	}
}
