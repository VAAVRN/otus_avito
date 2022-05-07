package main

import (
	"bufio"
	"fmt"
	"golang.org/x/example/stringutil"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	b, _, err := r.ReadLine()
	if err != nil {
		panic(err)
	}
	fmt.Println(stringutil.Reverse(string(b)))
}
