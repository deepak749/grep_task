package main

import (
	"fmt"
	"module/grepImplementation/grep"
)

func main() {
	files := []string{"test.txt", "test3.txt", "test2.txt"}
	flag := []string{"-n"}
	pattern := "hello"
	for _, s := range grep.Search(pattern, flag, files) {
		fmt.Println(s)
	}
}

