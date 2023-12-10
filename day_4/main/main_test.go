package main

import (
	"math"
	"slices"
	"testing"
)

// TODO: Add more test cases.

var ex = []string{
	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
}
var exOutputTotalSum = 13
var exOutputTotalCards = 30

func TestSumOfScratchCardPoints(t *testing.T) {
	res := sumOfScratchCardPoints(ex)
	if res != exOutputTotalSum {
		t.Errorf("expected %d but got %d",
			exOutputTotalSum, res)
	}
}

func TestGetTotalNumberOfScratchCardsAfterWinningOtherCards(t *testing.T) {
	res := getTotalNumberOfScratchCardsAfterWinningOtherCards(ex)
	if res != exOutputTotalCards {
		t.Errorf("expected %d but got %d",
			exOutputTotalCards, res)
	}
}

func TestSplitStringIntoWinningAndPlayerNumbers(t *testing.T) {
	type testCase struct {
		input         string
		outputWinning []string
		outputPlayer  []string
	}

	testCases := []testCase{ // example input provided by puzzle
		{
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			[]string{"41", "48", "83", "86", "17"},
			[]string{"83", "86", "6", "31", "17", "9", "48", "53"},
		},
		{
			"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			[]string{"13", "32", "20", "16", "61"},
			[]string{"61", "30", "68", "82", "17", "32", "24", "19"},
		},
		{
			"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			[]string{"1", "21", "53", "59", "44"},
			[]string{"69", "82", "63", "72", "16", "21", "14", "1"},
		},
		{
			"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			[]string{"41", "92", "73", "84", "69"},
			[]string{"59", "84", "76", "51", "58", "5", "54", "83"},
		},
		{
			"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			[]string{"87", "83", "26", "28", "32"},
			[]string{"88", "30", "70", "12", "93", "22", "82", "36"},
		},
		{
			"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			[]string{"31", "18", "13", "56", "72"},
			[]string{"74", "77", "10", "23", "35", "67", "36", "11"},
		},
	}

	for i, tst := range testCases {
		winning, player := splitStringIntoWinningAndPlayerNumbers(tst.input)
		slices.Sort(winning)
		slices.Sort(tst.outputWinning)
		slices.Sort(player)
		slices.Sort(tst.outputPlayer)
		wDiff := slices.Equal(winning, tst.outputWinning)
		pDiff := slices.Equal(player, tst.outputPlayer)

		if !wDiff {
			t.Errorf("test %d: strings for winning hand are not the same", i)
		}
		if !pDiff {
			t.Errorf("test %d: strings for player hand are not the same", i)
		}
	}
}

func TestGetMatchingNumbersInSCString(t *testing.T) {
	type testCase struct {
		input  string
		output int
	}

	testCases := []testCase{ // example input provided by puzzle
		{
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			4},
		{
			"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			2},
		{
			"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			2},
		{
			"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			1},
		{
			"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			0},
		{
			"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			0},
	}

	for i, tst := range testCases {
		res := getMatchingNumbersInSCString(tst.input)
		if res != tst.output {
			t.Errorf("test %d: expected %d but got %d",
				i, tst.output, res)
		}
	}
}

func TestGetNumberOfMatchingNumbers(t *testing.T) {
	type testCase struct {
		inputOne []string
		inputTwo []string
		output   int
	}

	testCases := []testCase{ // example input provided by puzzle
		{
			[]string{"41", "48", "83", "86", "17"},
			[]string{"83", "86", "6", "31", "17", "9", "48", "53"},
			4},
		{
			[]string{"13", "32", "20", "16", "61"},
			[]string{"61", "30", "68", "82", "17", "32", "24", "19"},
			2},
		{
			[]string{"1", "21", "53", "59", "44"},
			[]string{"69", "82", "63", "72", "16", "21", "14", "1"},
			2},
		{
			[]string{"41", "92", "73", "84", "69"},
			[]string{"59", "84", "76", "51", "58", "5", "54", "83"},
			1},
		{
			[]string{"87", "83", "26", "28", "32"},
			[]string{"88", "30", "70", "12", "93", "22", "82", "36"},
			0},
		{
			[]string{"31", "18", "13", "56", "72"},
			[]string{"74", "77", "10", "23", "35", "67", "36", "11"},
			0},
	}

	for i, tst := range testCases {
		res := getNumberOfMatchingNumbers(tst.inputOne, tst.inputTwo)
		if res != tst.output {
			t.Errorf("test %d: expected %d but got %d",
				i, tst.output, res)
		}
	}
}

func TestCalculateScratchCardValue(t *testing.T) {
	type testCase struct {
		input  int
		output int
	}

	testCases := []testCase{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 4},
		{5648264289, int(math.Pow(2, 5648264289))},
	}

	for i, tst := range testCases {
		res := calculateScratchCardValue(tst.input)
		if res != tst.output {
			t.Errorf("test %d: expected %d but got %d",
				i, tst.output, res)
		}
	}
}
