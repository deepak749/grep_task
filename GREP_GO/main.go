package main

import (
	"fmt"
	"module/grepImplementation/grep"
)

func main() {
	files := []string{"test.txt", "test2.txt", "test3.txt"}
	pattern := "hello"
	flag := []string{"-x"}
	for _, s := range grep.Search(pattern, flag, files) {
		fmt.Println(s)
	}
}
