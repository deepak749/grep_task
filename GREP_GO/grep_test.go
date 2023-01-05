package main

import (
	"fmt"
	"module/grepImplementation/grep"
	"testing"
)

func TestReturnGREP(t *testing.T) {
	file := []string{"testing.txt"}
	flag := []string{""}
	//testcase 0": flag==""
	fmt.Println("Testcase 0")
	actualSlice0 := grep.Search("hello", flag, file)
	expectedSlice0 := []string{"byye bye hellow", "gffv rree gghgv as hello"}
	for i, _ := range actualSlice0 {
		if actualSlice0[i] != expectedSlice0[i] {
			t.Errorf("Expected String(%s) is not same as"+
				" actual string (%s)", expectedSlice0, actualSlice0)
		}
	}

	//testcase 1: when flag is "-n"
	fmt.Println("Testcase 1")
	flag[0] = "-n"

	actualSlice1 := grep.Search("hello", flag, file)
	expectedSlice1 := []string{"2:byye bye hellow", "5:gffv rree gghgv as hello"}
	for i, _ := range actualSlice1 {
		if actualSlice1[i] != expectedSlice1[i] {
			t.Errorf("Expected String(%s) is not same as"+
				" actual string (%s)", expectedSlice1, actualSlice1)
		}
	}
	//testcase 2: when flag is "-l"
	fmt.Println("Testcase 2")
	flag[0] = "-l"

	expectedSlice2 := []string{"testing.txt"}
	actualSlice2 := grep.Search("hello", flag, file)
	for i, _ := range actualSlice2 {
		if actualSlice2[i] != expectedSlice2[i] {
			t.Errorf("Expected String(%s) is not same as"+
				" actual string (%s)", expectedSlice2, actualSlice2)
		}
	}

	//testcase 3: when flag is "-i"
	fmt.Println("Testcase 3")
	flag[0] = "-i"
	actualSlice3 := grep.Search("HELLO", flag, file)
	expectedSlice3 := []string{"byye bye hellow", "gffv rree gghgv as hello"}
	for i, _ := range actualSlice3 {
		if actualSlice3[i] != expectedSlice3[i] {
			t.Errorf("Expected String(%s) is not same as"+
				" actual string (%s)", expectedSlice3, actualSlice3)
		}
	}
	//testcase 4: when flag is "-v"
	fmt.Println("TestCase 4")
	flag[0] = "-v"
	expectedSlice4 := []string{"good bye heoll", "bohf", "dfdswfv"}
	actualSlice4 := grep.Search("hello", flag, file)
	for i, _ := range actualSlice4 {
		if actualSlice4[i] != expectedSlice4[i] {
			t.Errorf("Expected String(%s) is not same as"+
				" actual string (%s)", expectedSlice4, actualSlice4)
		}
	}

	//testcase 5: when flag is "-x"
	fmt.Println("TestCase 5")
	flag[0] = "-x"
	expectedSlice5 := []string{"gffv rree gghgv as hello"}
	actualSlice5 := grep.Search("hello", flag, file)
	for i, _ := range actualSlice5 {
		if actualSlice5[i] != expectedSlice5[i] {
			t.Errorf("Expected String(%s) is not same as"+
				" actual string (%s)", expectedSlice5, actualSlice5)
		}
	}

}
