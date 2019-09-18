package main

import "testing"

func TestRemoveWhitespace(t *testing.T) {
	input := " hello world \n \n   "
	expectedOutput := "helloworld"

	output := removeWhitespace(input)

	if output != expectedOutput {
		t.Errorf("Expected output of '%s', but got '%s'", expectedOutput, output)
	}
}
