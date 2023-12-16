package main

import (
	"slices"
	"testing"
)

func TestFillInLists(t *testing.T) {
	testInput := []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}

	seedToSoilOutput := []mappingStruct{
		{
			startValue:                     98,
			numberOfIntegersFollowingStart: 2,
			ouputValue:                     50,
		},
		{
			startValue:                     50,
			numberOfIntegersFollowingStart: 48,
			ouputValue:                     52,
		},
	}
	soilToFertilizerOutput := []mappingStruct{
		{
			startValue:                     15,
			numberOfIntegersFollowingStart: 37,
			ouputValue:                     0,
		},
		{
			startValue:                     52,
			numberOfIntegersFollowingStart: 2,
			ouputValue:                     37,
		},
		{
			startValue:                     0,
			numberOfIntegersFollowingStart: 15,
			ouputValue:                     39,
		},
	}
	fertilizerToWaterOutput := []mappingStruct{
		{
			startValue:                     53,
			numberOfIntegersFollowingStart: 8,
			ouputValue:                     49,
		},
		{
			startValue:                     11,
			numberOfIntegersFollowingStart: 42,
			ouputValue:                     0,
		},
		{
			startValue:                     0,
			numberOfIntegersFollowingStart: 7,
			ouputValue:                     42,
		},
		{
			startValue:                     7,
			numberOfIntegersFollowingStart: 4,
			ouputValue:                     57,
		},
	}
	waterToLightOutput := []mappingStruct{
		{
			startValue:                     18,
			numberOfIntegersFollowingStart: 7,
			ouputValue:                     88,
		},
		{
			startValue:                     25,
			numberOfIntegersFollowingStart: 70,
			ouputValue:                     18,
		},
	}
	lightToTemperatureOutput := []mappingStruct{
		{
			startValue:                     77,
			numberOfIntegersFollowingStart: 23,
			ouputValue:                     45,
		},
		{
			startValue:                     45,
			numberOfIntegersFollowingStart: 19,
			ouputValue:                     81,
		},
		{
			startValue:                     64,
			numberOfIntegersFollowingStart: 13,
			ouputValue:                     68,
		},
	}
	temperatureToHumidityOutput := []mappingStruct{
		{
			startValue:                     69,
			numberOfIntegersFollowingStart: 1,
			ouputValue:                     0,
		},
		{
			startValue:                     0,
			numberOfIntegersFollowingStart: 69,
			ouputValue:                     1,
		},
	}
	humidityToLocationOutput := []mappingStruct{
		{
			startValue:                     56,
			numberOfIntegersFollowingStart: 37,
			ouputValue:                     60,
		},
		{
			startValue:                     93,
			numberOfIntegersFollowingStart: 4,
			ouputValue:                     56,
		},
	}

	fillInLists(testInput)

	for _, v := range []int{79, 14, 55, 13} {
		if !slices.Contains(seeds, v) {
			t.Errorf("seed list does not contain seed %d", v)
		}
	}
	for _, v := range seedToSoilOutput {
		if !slices.Contains(seedToSoilOutput, v) {
			t.Errorf("seed to soil list does not contain start %d, step %d, output %d",
				v.startValue, v.numberOfIntegersFollowingStart, v.ouputValue)
		}
	}
	for _, v := range soilToFertilizerOutput {
		if !slices.Contains(soilToFertilizerOutput, v) {
			t.Errorf("soil to fertilizer list does not contain start %d, step %d, output %d",
				v.startValue, v.numberOfIntegersFollowingStart, v.ouputValue)
		}
	}
	for _, v := range fertilizerToWaterOutput {
		if !slices.Contains(fertilizerToWaterOutput, v) {
			t.Errorf("fertilizer to water list does not contain start %d, step %d, output %d",
				v.startValue, v.numberOfIntegersFollowingStart, v.ouputValue)
		}
	}
	for _, v := range waterToLightOutput {
		if !slices.Contains(waterToLightOutput, v) {
			t.Errorf("water to light list does not contain start %d, step %d, output %d",
				v.startValue, v.numberOfIntegersFollowingStart, v.ouputValue)
		}
	}
	for _, v := range lightToTemperatureOutput {
		if !slices.Contains(lightToTemperatureOutput, v) {
			t.Errorf("light to temperature list does not contain start %d, step %d, output %d",
				v.startValue, v.numberOfIntegersFollowingStart, v.ouputValue)
		}
	}
	for _, v := range temperatureToHumidityOutput {
		if !slices.Contains(temperatureToHumidityOutput, v) {
			t.Errorf("temperature to humidity list does not contain start %d, step %d, output %d",
				v.startValue, v.numberOfIntegersFollowingStart, v.ouputValue)
		}
	}
	for _, v := range humidityToLocationOutput {
		if !slices.Contains(humidityToLocationOutput, v) {
			t.Errorf("humidity to location list does not contain start %d, step %d, output %d",
				v.startValue, v.numberOfIntegersFollowingStart, v.ouputValue)
		}
	}

}

func TestPopulateSeedList(t *testing.T) {
	type testCase struct {
		input          string
		expectedOutput []int
	}

	testCases := []testCase{
		{
			"seeds: 79 14 55 13",
			[]int{79, 14, 55, 13},
		},
		{
			"seeds: 919339981 562444630 3366006921 67827214 1496677366 101156779 4140591657 5858311 2566406753 71724353 2721360939 35899538 383860877 424668759 3649554897 442182562 2846055542 49953829 2988140126 256306471",
			[]int{919339981, 562444630, 3366006921, 67827214, 1496677366, 101156779, 4140591657, 5858311, 2566406753, 71724353, 2721360939, 35899538, 383860877, 424668759, 3649554897, 442182562, 2846055542, 49953829, 2988140126, 256306471},
		},
	}

	for i, tst := range testCases {
		seeds = []int{}
		populateSeedList(tst.input)
		for _, v := range tst.expectedOutput {
			if !slices.Contains(seeds, v) {
				t.Errorf("test %d: seed list does not contain seed value: %d",
					i, v)
			}
		}
	}
}

func TestPopulateDataList(t *testing.T) {
	type testCase struct {
		input          []string
		mapToCheck     *[]mappingStruct
		expectedOutput []mappingStruct
	}

	testCases := []testCase{
		{
			input: []string{
				//"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
			},
			mapToCheck: &seedToSoils,
			expectedOutput: []mappingStruct{
				{
					startValue:                     98,
					numberOfIntegersFollowingStart: 2,
					ouputValue:                     50,
				},
				{
					startValue:                     50,
					numberOfIntegersFollowingStart: 48,
					ouputValue:                     52,
				},
			},
		},

		{
			input: []string{
				//"soil-to-fertilizer map:",
				"0 15 37",
				"37 52 2",
				"39 0 15",
			},
			mapToCheck: &soilToFertilizer,
			expectedOutput: []mappingStruct{
				{
					startValue:                     15,
					numberOfIntegersFollowingStart: 37,
					ouputValue:                     0,
				},
				{
					startValue:                     52,
					numberOfIntegersFollowingStart: 2,
					ouputValue:                     37,
				},
				{
					startValue:                     0,
					numberOfIntegersFollowingStart: 15,
					ouputValue:                     39,
				},
			},
		},
		{
			input: []string{
				//"fertilizer-to-water map:",
				"49 53 8",
				"0 11 42",
				"42 0 7",
				"57 7 4",
			},
			mapToCheck: &fertilizerToWater,
			expectedOutput: []mappingStruct{
				{
					startValue:                     53,
					numberOfIntegersFollowingStart: 8,
					ouputValue:                     49,
				},
				{
					startValue:                     11,
					numberOfIntegersFollowingStart: 42,
					ouputValue:                     0,
				},
				{
					startValue:                     0,
					numberOfIntegersFollowingStart: 7,
					ouputValue:                     42,
				},
				{
					startValue:                     7,
					numberOfIntegersFollowingStart: 4,
					ouputValue:                     57,
				},
			},
		},
		{
			input: []string{
				//"water-to-light map:",
				"88 18 7",
				"18 25 70",
			},
			mapToCheck: &waterToLight,
			expectedOutput: []mappingStruct{
				{
					startValue:                     18,
					numberOfIntegersFollowingStart: 7,
					ouputValue:                     88,
				},
				{
					startValue:                     25,
					numberOfIntegersFollowingStart: 70,
					ouputValue:                     18,
				},
			},
		},
		{
			input: []string{
				//"light-to-temperature map:",
				"45 77 23",
				"81 45 19",
				"68 64 13",
			},
			mapToCheck: &lightToTemperature,
			expectedOutput: []mappingStruct{
				{
					startValue:                     77,
					numberOfIntegersFollowingStart: 23,
					ouputValue:                     45,
				},
				{
					startValue:                     45,
					numberOfIntegersFollowingStart: 19,
					ouputValue:                     81,
				},
				{
					startValue:                     64,
					numberOfIntegersFollowingStart: 13,
					ouputValue:                     68,
				},
			},
		},
		{
			input: []string{
				//"temperature-to-humidity map:",
				"0 69 1",
				"1 0 69",
			},
			mapToCheck: &temperatureToHumidity,
			expectedOutput: []mappingStruct{
				{
					startValue:                     69,
					numberOfIntegersFollowingStart: 1,
					ouputValue:                     0,
				},
				{
					startValue:                     0,
					numberOfIntegersFollowingStart: 69,
					ouputValue:                     1,
				},
			},
		},
		{
			input: []string{
				//"humidity-to-location map:",
				"60 56 37",
				"56 93 4",
			},
			mapToCheck: &humidityToLocation,
			expectedOutput: []mappingStruct{
				{
					startValue:                     56,
					numberOfIntegersFollowingStart: 37,
					ouputValue:                     60,
				},
				{
					startValue:                     93,
					numberOfIntegersFollowingStart: 4,
					ouputValue:                     56,
				},
			},
		},
	}

	for i, tst := range testCases {
		populateDataList(tst.mapToCheck, tst.input)
		for _, v := range tst.expectedOutput {
			if !slices.Contains(*tst.mapToCheck, v) {
				t.Errorf("test %d: map did not contain correct value", i)
			}
		}
	}
}

func TestLookupMatchingTargetValueFromMappingStruct(t *testing.T) {
	type testCase struct {
		input          int
		lookUpList     *[]mappingStruct
		expectedOutput int
	}

	seedToSoilOutputLocal := []mappingStruct{
		{
			startValue:                     98,
			numberOfIntegersFollowingStart: 2,
			ouputValue:                     50,
		},
		{
			startValue:                     50,
			numberOfIntegersFollowingStart: 48,
			ouputValue:                     52,
		},
	}
	soilToFertilizerOutputLocal := []mappingStruct{
		{
			startValue:                     15,
			numberOfIntegersFollowingStart: 37,
			ouputValue:                     0,
		},
		{
			startValue:                     52,
			numberOfIntegersFollowingStart: 2,
			ouputValue:                     37,
		},
		{
			startValue:                     0,
			numberOfIntegersFollowingStart: 15,
			ouputValue:                     39,
		},
	}
	fertilizerToWaterOutputLocal := []mappingStruct{
		{
			startValue:                     53,
			numberOfIntegersFollowingStart: 8,
			ouputValue:                     49,
		},
		{
			startValue:                     11,
			numberOfIntegersFollowingStart: 42,
			ouputValue:                     0,
		},
		{
			startValue:                     0,
			numberOfIntegersFollowingStart: 7,
			ouputValue:                     42,
		},
		{
			startValue:                     7,
			numberOfIntegersFollowingStart: 4,
			ouputValue:                     57,
		},
	}
	waterToLightOutputLocal := []mappingStruct{
		{
			startValue:                     18,
			numberOfIntegersFollowingStart: 7,
			ouputValue:                     88,
		},
		{
			startValue:                     25,
			numberOfIntegersFollowingStart: 70,
			ouputValue:                     18,
		},
	}
	lightToTemperatureOutputLocal := []mappingStruct{
		{
			startValue:                     77,
			numberOfIntegersFollowingStart: 23,
			ouputValue:                     45,
		},
		{
			startValue:                     45,
			numberOfIntegersFollowingStart: 19,
			ouputValue:                     81,
		},
		{
			startValue:                     64,
			numberOfIntegersFollowingStart: 13,
			ouputValue:                     68,
		},
	}
	temperatureToHumidityOutputLocal := []mappingStruct{
		{
			startValue:                     69,
			numberOfIntegersFollowingStart: 1,
			ouputValue:                     0,
		},
		{
			startValue:                     0,
			numberOfIntegersFollowingStart: 69,
			ouputValue:                     1,
		},
	}
	humidityToLocationOutputLocal := []mappingStruct{
		{
			startValue:                     56,
			numberOfIntegersFollowingStart: 37,
			ouputValue:                     60,
		},
		{
			startValue:                     93,
			numberOfIntegersFollowingStart: 4,
			ouputValue:                     56,
		},
	}

	testCases := []testCase{
		{
			input:          79,
			lookUpList:     &seedToSoilOutputLocal,
			expectedOutput: 81,
		},
		{
			input:          81,
			lookUpList:     &soilToFertilizerOutputLocal,
			expectedOutput: 81,
		},
		{
			input:          81,
			lookUpList:     &fertilizerToWaterOutputLocal,
			expectedOutput: 81,
		},
		{
			input:          81,
			lookUpList:     &waterToLightOutputLocal,
			expectedOutput: 74,
		},
		{
			input:          74,
			lookUpList:     &lightToTemperatureOutputLocal,
			expectedOutput: 78,
		},
		{
			input:          78,
			lookUpList:     &temperatureToHumidityOutputLocal,
			expectedOutput: 78,
		},
		{
			input:          78,
			lookUpList:     &humidityToLocationOutputLocal,
			expectedOutput: 82,
		},
		{
			input:          14,
			lookUpList:     &seedToSoilOutputLocal,
			expectedOutput: 14,
		},
		{
			input:          14,
			lookUpList:     &soilToFertilizerOutputLocal,
			expectedOutput: 53,
		},
		{
			input:          53,
			lookUpList:     &fertilizerToWaterOutputLocal,
			expectedOutput: 49,
		},
		{
			input:          49,
			lookUpList:     &waterToLightOutputLocal,
			expectedOutput: 42,
		},
		{
			input:          42,
			lookUpList:     &lightToTemperatureOutputLocal,
			expectedOutput: 42,
		},
		{
			input:          42,
			lookUpList:     &temperatureToHumidityOutputLocal,
			expectedOutput: 43,
		},
		{
			input:          43,
			lookUpList:     &humidityToLocationOutputLocal,
			expectedOutput: 43,
		},

		{
			input:          55,
			lookUpList:     &seedToSoilOutputLocal,
			expectedOutput: 57,
		},
		{
			input:          57,
			lookUpList:     &soilToFertilizerOutputLocal,
			expectedOutput: 57,
		},
		{
			input:          57,
			lookUpList:     &fertilizerToWaterOutputLocal,
			expectedOutput: 53,
		},
		{
			input:          53,
			lookUpList:     &waterToLightOutputLocal,
			expectedOutput: 46,
		},
		{
			input:          46,
			lookUpList:     &lightToTemperatureOutputLocal,
			expectedOutput: 82,
		},
		{
			input:          82,
			lookUpList:     &temperatureToHumidityOutputLocal,
			expectedOutput: 82,
		},
		{
			input:          82,
			lookUpList:     &humidityToLocationOutputLocal,
			expectedOutput: 86,
		},

		{
			input:          13,
			lookUpList:     &seedToSoilOutputLocal,
			expectedOutput: 13,
		},
		{
			input:          13,
			lookUpList:     &soilToFertilizerOutputLocal,
			expectedOutput: 52,
		},
		{
			input:          52,
			lookUpList:     &fertilizerToWaterOutputLocal,
			expectedOutput: 41,
		},
		{
			input:          41,
			lookUpList:     &waterToLightOutputLocal,
			expectedOutput: 34,
		},
		{
			input:          34,
			lookUpList:     &lightToTemperatureOutputLocal,
			expectedOutput: 34,
		},
		{
			input:          34,
			lookUpList:     &temperatureToHumidityOutputLocal,
			expectedOutput: 35,
		},
		{
			input:          35,
			lookUpList:     &humidityToLocationOutputLocal,
			expectedOutput: 35,
		},
	}

	for i, tst := range testCases {
		res := lookupMatchingTargetValueFromMappingStruct(tst.input, tst.lookUpList)
		if res != tst.expectedOutput {
			t.Errorf("test %d: input %d should have output %d but got %d",
				i, tst.input, tst.expectedOutput, res)
		}
	}
}

func TestFindLowestLandValueFromSeedList(t *testing.T) {
	seeds = []int{79, 14, 55, 13}

	seedToSoils = []mappingStruct{
		{
			startValue:                     98,
			numberOfIntegersFollowingStart: 2,
			ouputValue:                     50,
		},
		{
			startValue:                     50,
			numberOfIntegersFollowingStart: 48,
			ouputValue:                     52,
		},
	}

	soilToFertilizer = []mappingStruct{
		{
			startValue:                     15,
			numberOfIntegersFollowingStart: 37,
			ouputValue:                     0,
		},
		{
			startValue:                     52,
			numberOfIntegersFollowingStart: 2,
			ouputValue:                     37,
		},
		{
			startValue:                     0,
			numberOfIntegersFollowingStart: 15,
			ouputValue:                     39,
		},
	}
	fertilizerToWater = []mappingStruct{
		{
			startValue:                     53,
			numberOfIntegersFollowingStart: 8,
			ouputValue:                     49,
		},
		{
			startValue:                     11,
			numberOfIntegersFollowingStart: 42,
			ouputValue:                     0,
		},
		{
			startValue:                     0,
			numberOfIntegersFollowingStart: 7,
			ouputValue:                     42,
		},
		{
			startValue:                     7,
			numberOfIntegersFollowingStart: 4,
			ouputValue:                     57,
		},
	}
	waterToLight = []mappingStruct{
		{
			startValue:                     18,
			numberOfIntegersFollowingStart: 7,
			ouputValue:                     88,
		},
		{
			startValue:                     25,
			numberOfIntegersFollowingStart: 70,
			ouputValue:                     18,
		},
	}
	lightToTemperature = []mappingStruct{
		{
			startValue:                     77,
			numberOfIntegersFollowingStart: 23,
			ouputValue:                     45,
		},
		{
			startValue:                     45,
			numberOfIntegersFollowingStart: 19,
			ouputValue:                     81,
		},
		{
			startValue:                     64,
			numberOfIntegersFollowingStart: 13,
			ouputValue:                     68,
		},
	}
	temperatureToHumidity = []mappingStruct{
		{
			startValue:                     69,
			numberOfIntegersFollowingStart: 1,
			ouputValue:                     0,
		},
		{
			startValue:                     0,
			numberOfIntegersFollowingStart: 69,
			ouputValue:                     1,
		},
	}
	humidityToLocation = []mappingStruct{
		{
			startValue:                     56,
			numberOfIntegersFollowingStart: 37,
			ouputValue:                     60,
		},
		{
			startValue:                     93,
			numberOfIntegersFollowingStart: 4,
			ouputValue:                     56,
		},
	}

	res := findLowestLandValueFromSeedList()
	if res != 35 {
		t.Errorf("final land value is incorrect. Expected %d but got %d",
			35, res)
	}
}
