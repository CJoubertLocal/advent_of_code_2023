package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

/*
Part 1:
Each card has a set of winning numbers, and a set of numbers given to the player.
If there is one number in both sets, the player gets 1 point.
For each additional number in both sets, the player's score doubles.

If size(intersection) == 0 => 0
If size(intersection) == 1 => 1
If size(intersection) == 2 => 2
If size(intersection) == 3 => 4

Plan:
Have an 'if' to check for the first case.
Then, 2^(size(intersection) - 1).

From the example input and input file, there are always fewer winning numbers in each
game than numbers which belong to the player.

Part 2:
Going forwards,
If card 1 wins 1 copy each of cards 2 and 3,
and card 2 wins 1 copy of card 3,
and card 3 wins 0 copies of any other card (or is the end of the sequence),
then there is
	1 copy of card 1 (original)
	2 copies of card 2 (original + card 1 winning)
	4 copies of card 3 (original + card 1 winning + card 2 winning + card 2 winning)

One idea:
Keep a tally, increment the next values in the tally by the number of instances of the current
case each time a card is one.

C|T   C|T   C|T
-|-   -|-   -|-
1|1 > 1|1 > 1|1
2|1   2|2   2|2
3|1   3|2   3|4

2's value becomes 1 + 1*1
3's value becomes 1 + 1*1 + 2*2

This should have a running time of O(n) as it reads the inputs once and never needs to look at them again.
Storage is also quite controlled. It is a hash map with n keys, and n integers.
*/

func main() {
	scratchCardLines := readInTextFile("./input.txt")
	res := sumOfScratchCardPoints(scratchCardLines)
	fmt.Println(res)
	totalNum := getTotalNumberOfScratchCardsAfterWinningOtherCards(scratchCardLines)
	fmt.Println(totalNum)
}

func readInTextFile(pathToTextFile string) []string {
	var linesFromFile []string

	file, err := os.Open(pathToTextFile)
	if err != nil {
		log.Fatal("unable to open text file with path ", pathToTextFile, " err ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linesFromFile = append(linesFromFile, scanner.Text())
	}

	return linesFromFile
}

func sumOfScratchCardPoints(scStrings []string) int {
	var sum = 0

	for _, sc := range scStrings {
		matchCount := getMatchingNumbersInSCString(sc)
		sum += calculateScratchCardValue(matchCount)
	}

	return sum
}

func getTotalNumberOfScratchCardsAfterWinningOtherCards(scStrings []string) int {
	var runningTotal = len(scStrings)
	var cardFrequency = map[int]int{}

	for i := 0; i < len(scStrings); i++ {
		cardFrequency[i] = 1
	}

	for i, sc := range scStrings {
		matchCount := getMatchingNumbersInSCString(sc)
		for j := 1; j < matchCount+1; j++ {
			if j == len(scStrings)-1 {
				break
			}
			cardFrequency[i+j] += cardFrequency[i]
			runningTotal += cardFrequency[i]
		}
	}

	return runningTotal
}

func splitStringIntoWinningAndPlayerNumbers(sc string) ([]string, []string) {
	splitString := strings.Split(sc[strings.Index(sc, ":")+1:], "|")
	return strings.Fields(splitString[0]), strings.Fields(splitString[1])
}

func getMatchingNumbersInSCString(sc string) int {
	winningSet, playerSet := splitStringIntoWinningAndPlayerNumbers(sc)
	return getNumberOfMatchingNumbers(winningSet, playerSet)
}

func getNumberOfMatchingNumbers(winningSet []string, playerSet []string) int {
	matchCount := 0
	wSMap := make(map[string]int)
	for _, wN := range winningSet {
		wSMap[wN] = 0
	}

	for _, pN := range playerSet {
		_, ok := wSMap[pN]
		if ok {
			matchCount++
		}
	}

	return matchCount
}

func calculateScratchCardValue(v int) int {
	switch v {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return int(math.Pow(2, float64(v-1)))
	}
}
