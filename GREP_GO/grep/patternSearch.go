package grep

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func flagsetting(text []string, flag, pattern string, file string) []string {
	result := []string{}
	if flag == "-n" {
		for i, line := range text {
			if strings.Contains(line, pattern) {
				str := strconv.Itoa(i + 1)
				result = append(result, str+":"+line)
			}
		}
	}
	if flag == "-l" {
		temp := false
		for _, line := range text {
			if strings.Contains(line, pattern) {
				temp = true
				break
			}
		}
		if temp {
			result = append(result, file)
		}
	}
	if flag == "-i" {
		for _, line := range text {
			if strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {

				result = append(result, line)
			}
		}

	}
	if flag == "-v" {
		for _, line := range text {
			if !strings.Contains(line, pattern) {
				result = append(result, line)
			}
		}
	}
	if flag == "-x" {
		for _, line := range text {
			temporarySlice := strings.Split(line, " ")
			for _, match := range temporarySlice {
				if match == pattern {
					result = append(result, line)
					break
				}
			}
		}
	}
	return result
}

func Search(pattern string, flags, fileName []string) []string {
	flag := flags[0]
	result := []string{}

	for _, files := range fileName {
		file, _ := os.Open(files)
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var text []string

		for scanner.Scan() {
			text = append(text, scanner.Text())
		}
		file.Close()
		result = append(result, flagsetting(text, flag, pattern, files)...)

	}
	return result
}
