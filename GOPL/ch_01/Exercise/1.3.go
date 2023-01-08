// measure running time of inefficient and efficient versions of echo
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("Efficient version: %d micro s\n", time.Since(start).Microseconds())
	start = time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("Inefficient version: %d micro s\n", time.Since(start).Microseconds())
}
