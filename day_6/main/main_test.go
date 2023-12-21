package main

import (
	"slices"
	"testing"
)

func TestConcatLinesInAllStrings(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	expectedOutput := []string{
		"Time:71530",
		"Distance:940200",
	}

	res := concatLinesInAllStrings(input)

	for _, e := range expectedOutput {
		if !slices.Contains(res, e) {
			t.Errorf("result does not contain expected output of %s", e)
		}
	}

	for _, r := range res {
		if !slices.Contains(expectedOutput, r) {
			t.Errorf("result contains unexpected output of %s", r)
		}
	}
}

func TestCreateBoatStructsFromTwoStringArrays(t *testing.T) {
	type testBoatRaceString struct {
		input          []string
		expectedOutput []boatRace
	}

	testCases := []testBoatRaceString{
		{
			input: []string{
				"Time:      7  15   30",
				"Distance:  9  40  200",
			},
			expectedOutput: []boatRace{
				{
					timeForRace:    7,
					distanceToBeat: 9,
				},
				{
					timeForRace:    15,
					distanceToBeat: 40,
				},
				{
					timeForRace:    30,
					distanceToBeat: 200,
				},
			},
		},
	}

	for i, tst := range testCases {
		res := createBoatStructsFromTwoStringArrays(tst.input)

		for _, eOut := range tst.expectedOutput {
			if !slices.Contains(res, eOut) {
				t.Errorf("test %d: expected result to contain race with time %d and distance %d",
					i, eOut.timeForRace, eOut.distanceToBeat)
			}
		}

		for _, resOut := range res {
			if !slices.Contains(tst.expectedOutput, resOut) {
				t.Errorf("test %d: result contains unexpected race with time %d and distance %d",
					i, resOut.timeForRace, resOut.distanceToBeat)
			}
		}
	}
}

func TestGetNumberOfWaysToWinAllRaces(t *testing.T) {
	type boatRaceTest struct {
		input          []boatRace
		expectedOutput int
	}

	testCases := []boatRaceTest{
		{
			input: []boatRace{
				{
					timeForRace:    7,
					distanceToBeat: 9,
				},
				{
					timeForRace:    15,
					distanceToBeat: 40,
				},
				{
					timeForRace:    30,
					distanceToBeat: 200,
				},
			},
			expectedOutput: 4 * 8 * 9,
		},
	}

	for i, tst := range testCases {
		res := getNumberOfWaysToWinAllRaces(tst.input)
		if res != tst.expectedOutput {
			t.Errorf("test %d: expected %d but got %d",
				i, tst.expectedOutput, res)
		}
	}
}

func TestGetNumberOfWaysToWinBoatRace(t *testing.T) {
	type boatRaceTest struct {
		input          boatRace
		expectedOutput int
	}

	testCases := []boatRaceTest{
		{
			input: boatRace{
				timeForRace:    7,
				distanceToBeat: 9,
			},
			expectedOutput: 4,
		},
		{
			input: boatRace{
				timeForRace:    15,
				distanceToBeat: 40,
			},
			expectedOutput: 8,
		},
		{
			input: boatRace{
				timeForRace:    30,
				distanceToBeat: 200,
			},
			expectedOutput: 9,
		},
	}

	for i, tst := range testCases {
		res := getNumberOfWaysToWinBoatRace(tst.input)
		if res != tst.expectedOutput {
			t.Errorf("test %d: expected %d but got %d", i, tst.expectedOutput, res)
		}
	}
}
