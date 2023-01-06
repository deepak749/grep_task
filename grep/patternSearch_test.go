package grep

import (
	"fmt"
	"testing"
)

func compareSlice(actualSlice, expectedSlice []string, t *testing.T) {
	for i, _ := range actualSlice {
		if actualSlice[i] != expectedSlice[i] {
			t.Errorf("Expected String(%s) is not same as"+
				" actual string (%s)", expectedSlice, actualSlice)
		}
	}
}

// testcase 0": flag==""
func TestCaseNoFlag(t *testing.T) {
	expectedSlice0 := []string{"byye bye hellow", "gffv rree gghgv as hello"}
	compareSlice(Search("hello", []string{""}, []string{"test_file1.txt"}), expectedSlice0, t)
}

// testcase 1: when flag is "-n"
func TestCaseHyphen_nFlag(t *testing.T) {

	expectedSlice1 := []string{"2:byye bye hellow", "5:gffv rree gghgv as hello"}
	compareSlice(Search("hello", []string{"-n"}, []string{"test_file1.txt"}), expectedSlice1, t)

}

//testcase 2: when flag is "-l"

func TestCaseHyphen_lFlag(t *testing.T) {
	expectedSlice2 := []string{"test_file1.txt"}
	compareSlice(Search("hello", []string{"-l"}, []string{"test_file1.txt"}), expectedSlice2, t)

}

//testcase 3: when flag is "-i"

func TestCaseHyph_iFlag(t *testing.T) {
	fmt.Println("Testcase 3")
	expectedSlice3 := []string{"byye bye hellow", "gffv rree gghgv as hello"}
	compareSlice(Search("hELLo", []string{"-i"}, []string{"test_file1.txt"}), expectedSlice3, t)
}

//testcase 4: when flag is "-v"

func TestCaseHyph_vFlag(t *testing.T) {
	expectedSlice4 := []string{"good bye heoll", "bohf", "dfdswfv"}
	compareSlice(Search("hello", []string{"-v"}, []string{"test_file1.txt"}), expectedSlice4, t)

}

//testcase 5: when flag is "-x"

func TestCaseHyph_xFlag(t *testing.T) {
	expectedSlice5 := []string{"gffv rree gghgv as hello"}
	compareSlice(Search("hello", []string{"-x"}, []string{"test_file1.txt"}), expectedSlice5, t)
}

//Testcase 6: if flag is "-n"

func TestCaseHyph_nWithMultiFile(t *testing.T) {
	// multiple file
	expectedSlice6 := []string{"test_file1.txt:2:byye bye hellow", "test_file1.txt:5:gffv rree gghgv as hello", "test_file2.txt:3:say hello"}
	compareSlice(Search("hello", []string{"-n"}, []string{"test_file1.txt", "test_file2.txt"}), expectedSlice6, t)

}
