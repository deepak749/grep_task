// package grep

// import (
// 	"bufio"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func MultipleFlagSituation(text []string, pattern string, file string, mpp map[string]int) []string {
// 	val := []string{}
// 	if mpp["-n"] == 1 && mpp["-v"] == 1 && mpp["-i"] == 0 && mpp["-l"] == 0 && mpp["-x"] == 0 {
// 		for i, line := range text {
// 			if !strings.Contains(line, pattern) {
// 				str := strconv.Itoa(i + 1)
// 				val = append(val, str+":"+line)
// 			}
// 		}

// 	}
// 	if mpp["-n"] == 1 && mpp["-x"] == 1 && mpp["-v"] == 1 && mpp["-i"] == 0 && mpp["-l"] == 0 {
// 		for i, line := range text {
// 			if line != pattern {
// 				str := strconv.Itoa(i + 1)
// 				val = append(val, str+":"+line)
// 			}
// 		}
// 	}
// 	if mpp["-n"] == 1 && mpp["-x"] == 1 && mpp["-v"] == 0 && mpp["-i"] == 0 && mpp["-l"] == 0 {
// 		for i, line := range text {
// 			if line == pattern {
// 				str := strconv.Itoa(i + 1)
// 				val = append(val, str+":"+line)
// 			}
// 		}
// 	}
// 	if mpp["-i"] == 1 && mpp["-v"] == 1 && mpp["-n"] == 0 && mpp["-x"] == 0 && mpp["-l"] == 0 {
// 		for _, line := range text {
// 			if !strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {
// 				val = append(val, line)
// 			}
// 		}
// 	}

// 	if mpp["-l"] == 1 && mpp["-v"] == 0 && mpp["-n"] == 0 && mpp["-x"] == 1 && mpp["-i"] == 0 {
// 		temp := false
// 		for _, line := range text {
// 			if line == pattern {
// 				temp = true
// 				break
// 			}
// 		}
// 		if temp {
// 			val = append(val, file)
// 		}
// 	}
// 	return val

// }

// func NoFlagMultipleFile(text []string, pattern string, file string) []string {
// 	v := []string{}
// 	for _, line := range text {
// 		if strings.Contains(line, pattern) {
// 			v = append(v, file+":"+line)
// 		}
// 	}
// 	return v
// }

// func NoFlagSingleFile(text []string, pattern string, file string) []string {
// 	val := []string{}
// 	for _, line := range text {
// 		if strings.Contains(line, pattern) {
// 			val = append(val, line)
// 		}
// 	}
// 	return val
// }
// func SingleFlag(text []string, flag, pattern string, file string, n int) []string {
// 	result := []string{}
// 	if flag == "-n" && n == 1 {
// 		for i, line := range text {
// 			if strings.Contains(line, pattern) {
// 				str := strconv.Itoa(i + 1)
// 				result = append(result, str+":"+line)
// 			}
// 		}
// 	}
// 	if flag == "-n" && n > 1 {
// 		for i, line := range text {
// 			if strings.Contains(line, pattern) {
// 				str := strconv.Itoa(i + 1)
// 				result = append(result, file+":"+str+":"+line)
// 			}
// 		}
// 	}
// 	if flag == "-l" {
// 		temp := false
// 		for _, line := range text {
// 			if strings.Contains(line, pattern) {
// 				temp = true
// 				break
// 			}
// 		}
// 		if temp {
// 			result = append(result, file)
// 		}
// 	}
// 	if flag == "-i" {
// 		for _, line := range text {
// 			if strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {

// 				result = append(result, line)
// 			}
// 		}

// 	}
// 	if flag == "-v" {
// 		for _, line := range text {
// 			if !strings.Contains(line, pattern) {
// 				result = append(result, line)
// 			}
// 		}
// 	}
// 	if flag == "-x" {
// 		for _, line := range text {
// 			if line == pattern {
// 				result = append(result, line)
// 			}
// 		}
// 	}
// 	return result
// }

// func Search(pattern string, flags, fileName []string) []string {
// 	n := len(fileName)
// 	result := []string{}

// 	for _, files := range fileName {
// 		file, _ := os.Open(files)
// 		scanner := bufio.NewScanner(file)
// 		scanner.Split(bufio.ScanLines)
// 		var text []string

// 		for scanner.Scan() {
// 			text = append(text, scanner.Text())
// 		}
// 		file.Close()
// 		if len(flags) <= 1 {
// 			if flags[0] != "" {
// 				result = append(result, SingleFlag(text, flags[0], pattern, files, n)...)

// 			} else {
// 				if n == 1 {
// 					result = append(result, NoFlagSingleFile(text, pattern, files)...)
// 				} else {
// 					result = append(result, NoFlagMultipleFile(text, pattern, files)...)

// 				}
// 			}
// 		} else {
// 			mpp := map[string]int{
// 				"-n": 0,
// 				"-l": 0,
// 				"-i": 0,
// 				"-v": 0,
// 				"-x": 0,
// 			}
// 			for _, flag := range flags {
// 				mpp[flag]++
// 			}
// 			result = append(result, MultipleFlagSituation(text, pattern, files, mpp)...)
// 		}

//		}
//		return result
//	}
package grep

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func MultipleFlagSituation(text []string, pattern string, file string, mpp map[string]int) []string {
	val := []string{}
	if mpp["-n"] == 1 && mpp["-v"] == 1 && mpp["-i"] == 0 && mpp["-l"] == 0 && mpp["-x"] == 0 {
		for i, line := range text {
			if !strings.Contains(line, pattern) {
				str := strconv.Itoa(i + 1)
				val = append(val, str+":"+line)
			}
		}

	}
	if mpp["-n"] == 1 && mpp["-x"] == 1 && mpp["-v"] == 1 && mpp["-i"] == 0 && mpp["-l"] == 0 {
		for i, line := range text {
			if line != pattern {
				str := strconv.Itoa(i + 1)
				val = append(val, str+":"+line)
			}
		}
	}
	if mpp["-n"] == 1 && mpp["-x"] == 1 && mpp["-v"] == 0 && mpp["-i"] == 0 && mpp["-l"] == 0 {
		for i, line := range text {
			if line == pattern {
				str := strconv.Itoa(i + 1)
				val = append(val, str+":"+line)
			}
		}
	}
	if mpp["-i"] == 1 && mpp["-v"] == 1 && mpp["-n"] == 0 && mpp["-x"] == 0 && mpp["-l"] == 0 {
		for _, line := range text {
			if !strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {
				val = append(val, line)
			}
		}
	}

	if mpp["-l"] == 1 && mpp["-v"] == 0 && mpp["-n"] == 0 && mpp["-x"] == 1 && mpp["-i"] == 0 {
		temp := false
		for _, line := range text {
			if line == pattern {
				temp = true
				break
			}
		}
		if temp {
			val = append(val, file)
		}
	}
	return val

}

func NoFlagMultipleFile(text []string, pattern string, file string) []string {
	v := []string{}
	for _, line := range text {
		if strings.Contains(line, pattern) {
			v = append(v, file+":"+line)
		}
	}
	return v
}

func NoFlagSingleFile(text []string, pattern string, file string) []string {
	val := []string{}
	for _, line := range text {
		if strings.Contains(line, pattern) {
			val = append(val, line)
		}
	}
	return val
}
func SingleFlag(text []string, flag, pattern string, file string, n int) []string {
	result := []string{}
	if flag == "-n" && n == 1 {
		for i, line := range text {
			if strings.Contains(line, pattern) {
				str := strconv.Itoa(i + 1)
				result = append(result, str+":"+line)
			}
		}
	}
	if flag == "-n" && n > 1 {
		for i, line := range text {
			if strings.Contains(line, pattern) {
				str := strconv.Itoa(i + 1)
				result = append(result, file+":"+str+":"+line)
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
			if line == pattern {
				result = append(result, line)
			}
		}
	}
	return result
}

func Search(pattern string, flags, fileName []string) []string {
	n := len(fileName)
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
		if len(flags) <= 1 {
			if flags[0] != "" {
				result = append(result, SingleFlag(text, flags[0], pattern, files, n)...)

			} else {
				if n == 1 {
					result = append(result, NoFlagSingleFile(text, pattern, files)...)
				} else {
					result = append(result, NoFlagMultipleFile(text, pattern, files)...)

				}
			}
		} else {
			mpp := map[string]int{
				"-n": 0,
				"-l": 0,
				"-i": 0,
				"-v": 0,
				"-x": 0,
			}
			for _, flag := range flags {
				mpp[flag]++
			}
			result = append(result, MultipleFlagSituation(text, pattern, files, mpp)...)
		}

	}
	return result
}
