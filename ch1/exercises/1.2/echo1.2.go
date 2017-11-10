// echo prints its command-line arguments.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i] + " "+ strconv.Itoa(i)
		sep = "\n"
	}
	fmt.Println(s)
}
