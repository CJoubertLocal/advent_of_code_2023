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

func TestSumOfPowersOfGames(t *testing.T) {
	type testCase struct {
		input          []string
		expectedOutput int
	}

	testCases := []testCase{
		{
			[]string{
				"Game 1: 0 red",
				"Game 234: 1 red",
			},
			1,
		},
		{
			[]string{
				"Game 43: 2 red",
				"Game 564: 2 red, 1 green"},
			4,
		},
		{
			[]string{
				"Game 23: 2 red, 1 green, 1 blue",
				"Game 654: 2 red, 1 green, 0 blue",
			},
			2,
		},
		{
			[]string{
				"Game 6: 20 red, 12 green, 0 blue",
				"Game 9876: 20 red, 12 green, 0 blue",
				"Game 73: 2 red; 1 green; 0 blue",
				"Game 29: 2 red; 1 green, 1  blue; 0 blue",
				"Game 29: 20 red, 47 green, 13 blue; 100 green, 46 blue; 46 green, 0 blue",
			},
			0 + 0 + 0 + 2 + 92000,
		},
		{
			[]string{
				"Game 76: 7 blue; 15 blue, 2 red, 1 green; 2 blue, 5 red; 2 red, 15 blue; 4 red, 15 blue; 9 blue, 5 red",
				"Game 77: 12 blue, 8 green, 15 red; 12 blue, 19 red, 16 green; 6 blue, 5 green, 16 red",
				"Game 78: 2 red, 7 blue, 14 green; 1 red, 3 green, 1 blue; 4 blue, 8 green, 10 red",
				"Game 79: 3 red, 5 green; 2 blue, 1 red, 18 green; 4 red, 15 green, 2 blue; 18 green, 2 blue, 7 red; 7 green, 6 red",
				"Game 80: 8 green, 5 red, 9 blue; 14 blue, 13 red, 6 green; 14 blue, 7 red, 4 green; 3 blue, 16 red, 4 green; 5 green, 13 blue, 2 red; 16 blue, 2 green, 5 red",
			},
			(5 * 1 * 15) + (19 * 16 * 12) + (10 * 14 * 7) + (7 * 18 * 2) + (16 * 8 * 16),
		},
	}

	for i, tst := range testCases {
		res := sumOfPowersOfGames(tst.input)
		if res != tst.expectedOutput {
			t.Errorf("game %d: input %s: expected %d, but got %d",
				i, tst.input, tst.expectedOutput, res)
		}
	}
}

func TestGetPowerOfGame(t *testing.T) {
	type testCase struct {
		input          string
		expectedOutput int
	}

	testCases := []testCase{
		{"Game 1: 0 red", 0},
		{"Game 234: 1 red", 1},
		{"Game 43: 2 red", 2},
		{"Game 564: 2 red, 1 green", 2},
		{"Game 23: 2 red, 1 green, 1 blue", 2},
		{"Game 654: 2 red, 1 green, 0 blue", 0},
		{"Game 6: 20 red, 7851 green, 0 blue", 0},
		{"Game 9876: 20 red, 7851 green, 0 blue", 0},
		{"Game 73: 2 red; 1 green; 0 blue", 0},
		{"Game 29: 2 red; 1 green, 1  blue; 0 blue", 2},
		{"Game 29: 20 red, 47 green, 13 blue; 100 green, 46 blue; 46 green, 0 blue", 92000},
	}

	for i, tst := range testCases {
		res := getPowerOfGame(tst.input)
		if res != tst.expectedOutput {
			t.Errorf("game %d: input %s: expected %d, but got %d",
				i, tst.input, tst.expectedOutput, res)
		}
	}
}
