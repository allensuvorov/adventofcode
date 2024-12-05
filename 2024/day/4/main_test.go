package main_test

import (
	"testing"
)

func TestCountXmas(t *testing.T) {
	// Test case 1: Empty matrix
	emptyMatrix := [][]byte{}
	expectedEmpty := 0
	if got := main.countXmas(emptyMatrix); got != expectedEmpty {
		t.Errorf("countXmas(emptyMatrix) expected %d, got %d", expectedEmpty, got)
	}

	// Test case 2: Small matrix with one "XMAS"
	testMatrix := [][]byte{
		{'X', 'M', 'A', 'S'},
		{'A', 'B', 'C', 'D'},
	}
	expectedOne := 1
	if got := main.countXmas(testMatrix); got != expectedOne {
		t.Errorf("countXmas(testMatrix) expected %d, got %d", expectedOne, got)
	}

	// Test case 3: Larger matrix with multiple occurrences
	largerMatrix := [][]byte{
		{'X', 'M', 'A', 'S'},
		{'A', 'B', 'X', 'S'},
		{'C', 'D', 'A', 'M'},
	}
	expectedMultiple := 3
	if got := main.countXmas(largerMatrix); got != expectedMultiple {
		t.Errorf("countXmas(largerMatrix) expected %d, got %d", expectedMultiple, got)
	}

	// You can add more test cases here for different scenarios
}
