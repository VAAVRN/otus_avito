package main

import (
	"bufio"
	"fmt"
	"golang.org/x/example/stringutil"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	b, _, _ := r.ReadLine()
	fmt.Println(stringutil.Reverse(string(b)))
}
