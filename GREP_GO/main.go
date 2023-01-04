package main

import (
	"bufio"
	"fmt"
	"module/grepImplementation/grep"
	"os"
	"strings"
)

func main() {
	files := []string{}
	flag := []string{}
	var pattern string
	fmt.Print("grep ")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')

	fmt.Println(userInput)
	inputs := strings.Split(userInput, " ")
	for _, param := range inputs {
		if strings.Contains(param, "-") {
			flag = append(flag, param)
		} else if strings.Contains(param, ".txt") {
			files = append(files, param)
		} else {
			pattern = param
		}
	}
	fmt.Println(files)
	for _, s := range grep.Search(pattern, flag, files) {
		fmt.Println(s)
	}
}
