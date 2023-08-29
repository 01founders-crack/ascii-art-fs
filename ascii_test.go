package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestPrintAllStringASCII(t *testing.T) {
	testCases := []struct {
		inputFile          string
		expectedOutputFile string
	}{
		{"testdata/input1.txt", "testdata/expected_output1.txt"},
		{"testdata/input2.txt", "testdata/expected_output2.txt"},
	}

	for _, tc := range testCases {
		// Load input and expected output files
		input, _ := ioutil.ReadFile(tc.inputFile)
		expectedOutput, _ := ioutil.ReadFile(tc.expectedOutputFile)

		fileLines := ReadStandardTxt()
		asciiTemplates := return2dASCIIArray(fileLines)

		// Call the function
		output := capturePrint(func() {
			printAllStringASCII(string(input), asciiTemplates)
		})

		// Compare the actual and expected output
		if strings.TrimSpace(output) != strings.TrimSpace(string(expectedOutput)) {
			t.Errorf("Expected output does not match actual output for input %s", tc.inputFile)
		}
	}
}

// Utility function to capture the printed output
func capturePrint(fn func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = old
	out, _ := ioutil.ReadAll(r)
	return string(out)
}
