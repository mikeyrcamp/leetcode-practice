package main

import (
	"fmt"
	"os"
)

func main() {
	s := []byte(os.Args[1])
	reverseString(s)
	fmt.Printf("%s\n", s)
}

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
