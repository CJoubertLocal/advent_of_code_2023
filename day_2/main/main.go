package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Code for part one.
TODO: Adapt for part two.
*/

func main() {
	linesFromFile := FindNumberOfPossibleGamesInTextFile("../input.txt")
	fmt.Println(sumAllGames(linesFromFile))
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

func sumAllGames(allGames []string) int {
	var runningTotal = 0

	for gameNumber, game := range allGames {
		if checkIfAValidGame(game) {
			runningTotal += gameNumber + 1
		}
		fmt.Println(gameNumber, runningTotal, game)
	}

	return runningTotal
}

func checkIfAValidGame(gameLine string) bool {
	hands := strings.Split(gameLine[strings.Index(gameLine, ":")+2:], ";")

	for _, hand := range hands {
		if !checkHand(hand) {
			return false
		}
	}

	return true
}

func checkHand(handLine string) bool {
	eachHand := strings.Split(handLine, ",")
	for _, hand := range eachHand {
		countAndColour := strings.Fields(hand)
		if colourCountIsGreaterThanMaximum(countAndColour[0], countAndColour[1]) {
			return false
		}
	}

	return true
}

var colourMaximums = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func colourCountIsGreaterThanMaximum(colourCount string, colour string) bool {
	num, err := strconv.Atoi(colourCount)
	if err != nil {
		log.Printf("error: unable to parse int for %s", colourCount)
	}

	if num > colourMaximums[colour] {
		return true
	}

	return false
}
