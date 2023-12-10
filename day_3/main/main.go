package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
Efficiencies to make:
	Replace looping over the input twice with a single pass.
	Only keep three lines in memory. This is the maximum needed to check for symbols / numbers adjacent to
		the symbols and numbers on the current line.
	If the relative frequency of symbols / numbers is known, then search for the one with a lower frequency
		and then search in the surrounding positions.
*/

type gearLoc struct {
	r int
	c int
}

var numberRunes = []uint8{
	uint8('0'),
	uint8('1'),
	uint8('2'),
	uint8('3'),
	uint8('4'),
	uint8('5'),
	uint8('6'),
	uint8('7'),
	uint8('8'),
	uint8('9'),
}

func main() {
	linesFromFile := FindNumberOfPossibleGamesInTextFile("../input.txt")

	// inefficient 2n approach - makes two passes
	// part 1
	numbersNextToSymbols := getNumbersNextToSymbols(linesFromFile)
	sum := 0
	for _, n := range numbersNextToSymbols {
		sum += n
	}
	fmt.Println(sum)

	// part 2
	fmt.Println(getSumOfGearNumberProducts(linesFromFile))
}

func FindNumberOfPossibleGamesInTextFile(pathToTextFile string) []string {
	var linesFromFile []string

	file, err := os.Open(pathToTextFile)
	if err != nil {
		log.Fatal("unable to open text file with path ", pathToTextFile)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linesFromFile = append(linesFromFile, scanner.Text())
	}

	return linesFromFile
}

func getNumbersNextToSymbols(textIn []string) []int {
	res := []int{}

	for i := 0; i < len(textIn); i++ {
		curStartingIndex := -1
		for j := 0; j < len(textIn[i]); j++ {
			if slices.Contains(numberRunes, textIn[i][j]) && curStartingIndex == -1 {
				curStartingIndex = j
			}

			if (!slices.Contains(numberRunes, textIn[i][j]) || (j == len(textIn[i])-1)) && curStartingIndex > -1 {

				endIdx := j
				if (j == len(textIn[i])-1) && slices.Contains(numberRunes, textIn[i][j]) {
					endIdx++
				}

				if symbolInDiagonal(i, curStartingIndex, endIdx, textIn) {
					n, err := strconv.Atoi(textIn[i][curStartingIndex:endIdx])
					if err != nil {
						log.Println(err)
					}
					res = append(res, n)
				}

				curStartingIndex = -1
			}
		}
	}

	return res
}

// Simplification needed
func symbolInDiagonal(rowNum int, startingCol int, endingCol int, textIn []string) bool {
	if rowNum < 0 || len(textIn) < rowNum || startingCol < 0 || len(textIn[rowNum]) < endingCol {
		return false
	}

	for i := rowNum - 1; i <= rowNum+1; i++ {
		if -1 < i && i < len(textIn) {
			for j := startingCol - 1; j < endingCol+1; j++ {
				if -1 < j && j < len(textIn[rowNum]) {
					if !slices.Contains(numberRunes, textIn[i][j]) && textIn[i][j] != uint8('.') { // part one
						return true
					}
				}
			}
		}
	}

	return false
}

func getSumOfGearNumberProducts(textIn []string) int {
	res := 0
	prods := getProductOfIntPairsNextToGears(textIn)

	for _, n := range prods {
		res += n
	}

	return res
}

func getProductOfIntPairsNextToGears(textIn []string) []int {
	askLocs := getAsteriskLocations(textIn)
	var res []int

	for _, l := range askLocs {
		var numsForAskterisk []int
		for i := l.r - 1; i <= l.r+1; i++ {

			if -1 < i && i < len(textIn) {
				ns := getNumbersInStringAroundIndex(textIn[i], l.c)

				for _, n := range ns {
					numsForAskterisk = append(numsForAskterisk, n)
				}
			}
		}

		if len(numsForAskterisk) == 2 {
			res = append(res, numsForAskterisk[0]*numsForAskterisk[1])
		}
	}

	return res
}

func getAsteriskLocations(textIn []string) []gearLoc {
	glocs := []gearLoc{}

	for i, l := range textIn {
		startIdx := 0

		numAsk := strings.Count(l, "*")
		for j := 0; j < numAsk; j++ {
			startIdx += strings.Index(l[startIdx:], "*")

			glocs = append(glocs, gearLoc{i, startIdx})
			startIdx++
		}
	}

	return glocs
}

func getNumbersInStringAroundIndex(str string, idx int) []int {
	var res []int

	startIdx := -1
	for i := 0; i < len(str); i++ {
		if slices.Contains(numberRunes, str[i]) && startIdx == -1 {
			startIdx = i
		}

		if (!slices.Contains(numberRunes, str[i]) || i == len(str)-1) && -1 < startIdx {
			endIdx := i
			if endIdx == len(str)-1 && slices.Contains(numberRunes, str[i]) {
				endIdx++
			}

			if idx == startIdx-1 || (startIdx <= idx && idx <= endIdx) || idx == endIdx {
				n, _ := strconv.Atoi(str[startIdx:endIdx])
				res = append(res, n)

				startIdx = -1
			}
		}

		if !slices.Contains(numberRunes, str[i]) {
			startIdx = -1
		}
	}

	return res
}
