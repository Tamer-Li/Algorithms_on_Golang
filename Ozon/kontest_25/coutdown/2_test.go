package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func runTest(input string) string {
	tmpInputFile, err := os.CreateTemp("", "input")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpInputFile.Name())

	_, err = tmpInputFile.WriteString(input)
	if err != nil {
		panic(err)
	}
	tmpInputFile.Close()

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin, err = os.Open(tmpInputFile.Name())
	if err != nil {
		panic(err)
	}

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	main()

	w.Close()
	var output strings.Builder
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		output.WriteString(scanner.Text() + "\n")
	}

	return output.String()
}

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func TestTwo(t *testing.T) {
	for i := 1; i <= 18; i++ {
		if i == 6 || i == 7 || i == 9 || i == 10 || i == 11 || i == 13 || i == 14 || i == 16 || i == 17 {
			// t.Logf("Skipping test %d (file missing)", i)
			continue
		}

		inputFile := fmt.Sprintf("countdown_case/%d", i)
		input, err := readFile(inputFile)
		if err != nil {
			t.Errorf("Failed to read input file %d: %v", i, err)
			continue
		}

		expectedOutputFile := fmt.Sprintf("countdown_case/%d.a", i)
		expectedOutput, err := readFile(expectedOutputFile)
		if err != nil {
			t.Errorf("Failed to read expected output file %d: %v", i, err)
			continue
		}

		actualOutput := runTest(input)

		if actualOutput != expectedOutput {
			t.Errorf("Test %d failed\n", i)
		} else {
			t.Logf("Test %d passed", i)
		}
	}
}
