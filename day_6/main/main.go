package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type boatRace struct {
	timeForRace    int
	distanceToBeat int
}

func main() {
	linesFromFile := readInTextFile("../input.txt")
	boatStructs := createBoatStructsFromTwoStringArrays(linesFromFile)
	fmt.Println(getNumberOfWaysToWinAllRaces(boatStructs))

	kerningRace := concatLinesInAllStrings(linesFromFile)
	oneRaceBoatStruct := createBoatStructsFromTwoStringArrays(kerningRace)
	fmt.Println(getNumberOfWaysToWinAllRaces(oneRaceBoatStruct))
}

func concatLinesInAllStrings(stringsIn []string) []string {
	var res []string

	for _, s := range stringsIn {
		res = append(res, strings.ReplaceAll(s, " ", ""))
	}

	return res
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

// createBoatStructsFromTwoStringArrays assumes that the input contains exactly
// two strings.
func createBoatStructsFromTwoStringArrays(linesIn []string) []boatRace {
	var bRaces = []boatRace{}

	raceTimes := strings.Fields(linesIn[0][strings.Index(linesIn[0], ":")+1:])
	distancesToBeat := strings.Fields(linesIn[1][strings.Index(linesIn[1], ":")+1:])

	for i := 0; i < len(raceTimes); i++ {
		tInt, err := strconv.Atoi(raceTimes[i])
		if err != nil {
			log.Printf("error when converting time %s to int\n", raceTimes[i])
			log.Println(err)
		}

		dInt, err := strconv.Atoi(distancesToBeat[i])
		if err != nil {
			log.Printf("error when converting distance %s to int\n", distancesToBeat[i])
			log.Println(err)
		}

		bRaces = append(bRaces, boatRace{
			timeForRace:    tInt,
			distanceToBeat: dInt,
		})
	}

	return bRaces
}

func getNumberOfWaysToWinAllRaces(boatRaces []boatRace) int {
	var sum = 1

	for i := 0; i < len(boatRaces); i++ {
		sum *= getNumberOfWaysToWinBoatRace(boatRaces[i])
	}

	return sum
}

func getNumberOfWaysToWinBoatRace(race boatRace) int {
	var sum = 0

	for i := 1; i < race.timeForRace; i++ {
		if i*(race.timeForRace-i) > race.distanceToBeat {
			sum++
		}
	}

	return sum
}
