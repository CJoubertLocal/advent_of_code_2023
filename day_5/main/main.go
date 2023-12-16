package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type mappingStruct struct {
	startValue                     int
	numberOfIntegersFollowingStart int
	ouputValue                     int
}

var seeds []int
var seedToSoils []mappingStruct
var soilToFertilizer []mappingStruct
var fertilizerToWater []mappingStruct
var waterToLight []mappingStruct
var lightToTemperature []mappingStruct
var temperatureToHumidity []mappingStruct
var humidityToLocation []mappingStruct

func main() {
	linesFromFile := readInTextFile("../input.txt")
	fillInLists(linesFromFile)
	fmt.Println(findLowestLandValueFromSeedList())
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

// TODO: This is inefficient. I would be neater to, whilst reading in the file, identify when a new section starts
// and then read until a line of length 0.
func fillInLists(stringsIn []string) {

	err := populateSeedList(stringsIn[0])
	if err != nil {
		log.Printf("issue in populating seed list")
	}

	// move to file reading for loop
	seedToSoilsStartIndex := -1
	seedToSoilsEndIndex := -1
	soilToFertilizerStartIndex := -1
	soilToFertilizerEndIndex := -1
	fertilizerToWaterStartIndex := -1
	fertilizerToWaterEndIndex := -1
	waterToLightStartIndex := -1
	waterToLightEndIndex := -1
	lightToTemperatureStartIndex := -1
	lightToTemperatureEndIndex := -1
	temperatureToHumidityStartIndex := -1
	temperatureToHumidityEndIndex := -1
	humidityToLocationStartIndex := -1
	humidityToLocationEndIndex := len(stringsIn) - 1

	for i := 1; i < len(stringsIn); i++ {
		if strings.Contains(stringsIn[i], "seed-to-soil map:") {
			seedToSoilsStartIndex = i + 1
		}
		if strings.Contains(stringsIn[i], "soil-to-fertilizer map:") {
			seedToSoilsEndIndex = i - 1
			soilToFertilizerStartIndex = i + 1
		}
		if strings.Contains(stringsIn[i], "fertilizer-to-water map:") {
			soilToFertilizerEndIndex = i - 1
			fertilizerToWaterStartIndex = i + 1
		}
		if strings.Contains(stringsIn[i], "water-to-light map:") {
			fertilizerToWaterEndIndex = i - 1
			waterToLightStartIndex = i + 1
		}
		if strings.Contains(stringsIn[i], "light-to-temperature map:") {
			waterToLightEndIndex = i - 1
			lightToTemperatureStartIndex = i + 1
		}
		if strings.Contains(stringsIn[i], "temperature-to-humidity map:") {
			lightToTemperatureEndIndex = i - 1
			temperatureToHumidityStartIndex = i + 1
		}
		if strings.Contains(stringsIn[i], "humidity-to-location map:") {
			temperatureToHumidityEndIndex = i - 1
			humidityToLocationStartIndex = i + 1
		}
	}

	populateDataList(&seedToSoils, stringsIn[seedToSoilsStartIndex:seedToSoilsEndIndex])
	populateDataList(&soilToFertilizer, stringsIn[soilToFertilizerStartIndex:soilToFertilizerEndIndex])
	populateDataList(&fertilizerToWater, stringsIn[fertilizerToWaterStartIndex:fertilizerToWaterEndIndex])
	populateDataList(&waterToLight, stringsIn[waterToLightStartIndex:waterToLightEndIndex])
	populateDataList(&lightToTemperature, stringsIn[lightToTemperatureStartIndex:lightToTemperatureEndIndex])
	populateDataList(&temperatureToHumidity, stringsIn[temperatureToHumidityStartIndex:temperatureToHumidityEndIndex])
	populateDataList(&humidityToLocation, stringsIn[humidityToLocationStartIndex:humidityToLocationEndIndex])

}

func populateSeedList(stringIn string) error {
	splitString := strings.Fields(stringIn)

	for i := 1; i < len(splitString); i++ {
		seedInt, err := strconv.Atoi(splitString[i])
		if err != nil {
			log.Printf("unable to convert string to int: %s", err)
			return err
		}
		seeds = append(seeds, seedInt)
	}

	return nil
}

func populateDataList(listToPopulate *[]mappingStruct, stringsIn []string) error {
	for i := 0; i < len(stringsIn); i++ {
		splitString := strings.Fields(stringsIn[i])

		targetStart, err := strconv.Atoi(splitString[0])
		if err != nil {
			log.Printf("unable to convert string to int: %s", err)
			return err
		}

		startVal, err := strconv.Atoi(splitString[1])
		if err != nil {
			log.Printf("unable to convert string to int: %s", err)
			return err
		}

		numberOfFollowingValues, err := strconv.Atoi(splitString[2])
		if err != nil {
			log.Printf("unable to convert string to int: %s", err)
			return err
		}

		*listToPopulate = append(*listToPopulate, mappingStruct{
			startValue:                     startVal,
			numberOfIntegersFollowingStart: numberOfFollowingValues,
			ouputValue:                     targetStart,
		})
	}

	return nil
}

func lookupMatchingTargetValueFromMappingStruct(valueOfInterest int, referenceStructList *[]mappingStruct) int {
	for _, m := range *referenceStructList {
		if m.startValue <= valueOfInterest && valueOfInterest < m.startValue+m.numberOfIntegersFollowingStart {
			return m.ouputValue + valueOfInterest - m.startValue
		}
	}

	return valueOfInterest
}

func findLowestLandValueFromSeedList() int {
	lowestLandVal := -1
	for _, s := range seeds {
		soil := lookupMatchingTargetValueFromMappingStruct(s, &seedToSoils)
		fert := lookupMatchingTargetValueFromMappingStruct(soil, &soilToFertilizer)
		water := lookupMatchingTargetValueFromMappingStruct(fert, &fertilizerToWater)
		light := lookupMatchingTargetValueFromMappingStruct(water, &waterToLight)
		temp := lookupMatchingTargetValueFromMappingStruct(light, &lightToTemperature)
		humid := lookupMatchingTargetValueFromMappingStruct(temp, &temperatureToHumidity)
		loc := lookupMatchingTargetValueFromMappingStruct(humid, &humidityToLocation)

		if lowestLandVal == -1 || loc < lowestLandVal {
			lowestLandVal = loc
		}
	}

	return lowestLandVal
}
