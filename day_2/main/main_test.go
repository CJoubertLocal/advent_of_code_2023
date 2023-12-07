package main

import (
	"testing"
)

// Tests for part one

func TestCheckIfAValidGame(t *testing.T) {
	type testCase struct {
		input          string
		expectedOutput bool
	}

	testCases := []testCase{
		{"1 red", true},
		{"1 green", true},
		{"1 blue", true},
		{"12 red", true},
		{"13 green", true},
		{"14 blue", true},
		{"13 red", false},
		{"14 green", false},
		{"15 blue", false},
		// from website
		{"3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", true},
		{"1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", true},
		{"8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", false},
		{"1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", false},
		{"6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", true},
	}

	for i, tst := range testCases {
		res := checkIfAValidGame(tst.input)
		if res != tst.expectedOutput {
			t.Errorf("test %d: input %s: expected output %t but got %t",
				i, tst.input, tst.expectedOutput, res)
		}
	}
}

func TestCheckHand(t *testing.T) {
	type testCase struct {
		input          string
		expectedOutput bool
	}

	testCases := []testCase{
		{"1 red", true},
		{"1 green", true},
		{"1 blue", true},
		{"12 red", true},
		{"13 green", true},
		{"14 blue", true},
		{"13 red", false},
		{"14 green", false},
		{"15 blue", false},
		// from website
		{"3 blue, 4 red", true},
		{"1 red, 2 green, 6 blue", true},
		{"2 green", true},
		{"1 blue, 2 green", true},
		{"3 green, 4 blue, 1 red", true},
		{"1 green, 1 blue", true},
		{"8 green, 6 blue, 20 red", false},
		{"5 blue, 4 red, 13 green", true},
		{"5 green, 1 red", true},
		{"1 green, 3 red, 6 blue", true},
		{"3 green, 6 red", true},
		{"3 green, 15 blue, 14 red", false},
		{"6 red, 1 blue, 3 green", true},
		{"2 blue, 1 red, 2 green", true},
	}

	for i, tst := range testCases {
		res := checkHand(tst.input)
		if res != tst.expectedOutput {
			t.Errorf("test %d: input %s: expected output %t but got %t",
				i, tst.input, tst.expectedOutput, res)
		}
	}
}

func TestColourCountIsGreaterThanMaximum(t *testing.T) {
	type testCase struct {
		numberInput    string
		colour         string
		expectedOutput bool
	}

	testCases := []testCase{
		{"1", "red", false},
		{"1", "green", false},
		{"1", "blue", false},
		{"12", "red", false},
		{"13", "green", false},
		{"14", "blue", false},
		{"13", "red", true},
		{"14", "green", true},
		{"15", "blue", true},
	}

	for i, tst := range testCases {
		res := colourCountIsGreaterThanMaximum(tst.numberInput, tst.colour)
		if res != tst.expectedOutput {
			t.Errorf("test %d: input (%s, %s): expected output %t but got %t",
				i, tst.numberInput, tst.colour, tst.expectedOutput, res)
		}
	}
}
