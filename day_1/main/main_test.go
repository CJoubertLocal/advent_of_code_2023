package main

import "testing"

func TestSumOfFirstAndLastDigitsFromTextStrings(t *testing.T) {
	type testCase struct {
		input          []string
		expectedOutput int
	}

	testCases := []testCase{
		{[]string{"00"}, 0},
		{[]string{"1"}, 11},
		{[]string{"123"}, 13},
		{[]string{"123", "456"}, 13 + 46},
		{[]string{"a1b2c3d", "e4f5g6hiforty"}, 13 + 46},
		{[]string{ // from input text file
			"heightwosixthzdf7gdtllhsnfive1onemfcqkqfqkj1",
			"429ninennkhtzveight7lqmzdlgg",
			"eightgfhrxtvcsg38",
		}, 81 + 47 + 88}, // digit only output: 71 + 47 + 38},
		{[]string{ // partly from input test file
			"4sixbcxmgvfourjsqkc85eight8",
			"2threerfjsgfdlhheightfive",
			"five5thirtyeight38",
		}, 48 + 25 + 58}, // digit only output: 48 + 22 + 58},
	}

	for i, tst := range testCases {
		res := SumOfFirstAndLastDigitsFromTextStrings(tst.input)
		if res != tst.expectedOutput {
			t.Errorf("test %d: input %s: expected %d, got %d",
				i, tst.input, tst.expectedOutput, res,
			)
		}
	}

}

func TestConcatenateDigitsInString(t *testing.T) {
	type testCase struct {
		input  string
		output int
	}

	testCases := []testCase{
		{"abc12cba", 12},
		// examples provided by advent of code
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
		// examples from input file provided by advent of code
		{"heightwosixthzdf7gdtllhsnfive1onemfcqkqfqkj1",
			81}, // digit only output: 71},
		{"twofoureightpqpscvhlvbxgfpbrrkvc6sixxzhn",
			26}, // digit only output: 66},
		{"chbfbjctpkdtmkjksct1ninetwoeight9vbnv",
			19}, // digit only output: 19},
		{"2vjnvdv",
			22},
		{"429ninennkhtzveight7lqmzdlgg",
			47},
		{"xqeightwothree9119qmhdltcmvdmxzvsix",
			86}, // digit only output: 99},
		{"eightgfhrxtvcsg38",
			88}, // digit only output: 38},
		{"twomhll7sxjtwoone7dhzbhphlpmlhx",
			27}, // digit only output: 77},
	}

	for i, tst := range testCases {
		res := ConcatenateDigitsInString(tst.input)
		if res != tst.output {
			t.Errorf("test %d: input %s: expected output %d but got output %d", i, tst.input, tst.output, res)
		}
	}
}

func TestFindDigitInString(t *testing.T) {
	type inAndOut struct {
		input          string
		searchFromEnd  bool
		expectedOutput string
	}

	inputsAndExpectedOutputs := []inAndOut{
		{"abc12cba", false, "1"},
		{"abc12cba", true, "2"},
		{"abc12cba3", true, "3"},
		{"abc142cba3", false, "1"},
		{"sevenabc142cba3", false, "7"},
		{"abc142cba3four", true, "4"},
		{"abc142cbafour3", true, "3"},
		{"onetwo34fivesix", false, "1"},
		{"onetwo34fivesix", true, "6"},
		{"oonetwo34fivesix", true, "6"},
		{"oonetwo34fivesix", false, "1"},
	}

	for i, tst := range inputsAndExpectedOutputs {
		res := FindDigitInString(tst.input, tst.searchFromEnd)
		if res != tst.expectedOutput {
			t.Errorf("test %d:string %s: expected %s but got %s", i, tst.input, tst.expectedOutput, res)
		}
	}
}

func TestCheckIfStringContainsDigitNameStartingAtPosition(t *testing.T) {
	type testCase struct {
		input          string
		index          int
		expectedOutput string
	}

	testCases := []testCase{
		{"one2", 0, "1"},
		{"one2", 1, ""},
		{"kone2", 1, "1"},
		{"one2", 2, ""},
		{"one2", 3, ""},
		{"one2three", 3, ""},
		{"one2three", 4, "3"},
	}

	for i, tst := range testCases {
		res := CheckIfStringContainsDigitNameStartingAtPosition(tst.input, tst.index)
		if res != tst.expectedOutput {
			t.Errorf("test %d: input %s: expected %s but got %s",
				i, tst.input, tst.expectedOutput, res,
			)
		}
	}
}
