package main

import (
	"bufio"
	"fmt"
	"module/grepImplementation/grep"
	"os"
	"strings"
)

func main() {

	flag := []string{}
	pattern := "hello"
	files := []string{}

	fmt.Print("grep ")
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	scanner.Scan()
	line := scanner.Text()
	lines = strings.Split(line, " ")
	fmt.Println(len(lines))
	pattern = lines[0]
	for i := 1; i < len(lines); i++ {
		if lines[i][0] == '-' {
			flag = append(flag, lines[i])
		} else {
			files = append(files, lines[i])
		}
	}
	if len(flag) == 0 {
		flag = append(flag, "")
	}
	for _, s := range grep.Search(pattern, flag, files) {
		fmt.Println(s)
	}

}
