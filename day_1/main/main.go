package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	tokensToRead := ReadInTextFile("../input.txt")
	fmt.Println(SumOfFirstAndLastDigitsFromTextStrings(tokensToRead))
}

func ReadInTextFile(pathToTextFile string) []string {
	// https://gobyexample.com/reading-files
	// https://zetcode.com/golang/readfile/
	var linesOfText []string

	file, err := os.Open(pathToTextFile)
	if err != nil {
		log.Fatal("unable to open text file with path ", pathToTextFile)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linesOfText = append(linesOfText, scanner.Text())
	}

	return linesOfText
}

func SumOfFirstAndLastDigitsFromTextStrings(strsIn []string) int {
	sum := 0

	for _, currentStr := range strsIn {
		sum += ConcatenateDigitsInString(currentStr)
	}

	return sum
}

func ConcatenateDigitsInString(strIn string) int {
	concatenatedString := FindDigitInString(strIn, false) + FindDigitInString(strIn, true)

	intout, err := strconv.Atoi(concatenatedString)
	if err != nil {
		log.Println("unable to convert string to int ", concatenatedString)
	}

	return intout
}

var digitStartingChars = []rune{
	'e', 'f', 'n', 'o', 's', 't',
}

var digitToIntMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func FindDigitInString(strIn string, startFromEnd bool) string {

	runesOfStr := []rune(strIn)

	if !startFromEnd {
		for i := 0; i < len(runesOfStr); i++ {
			res := helperForFindDigitInString(strIn, runesOfStr[i], i)
			if res != "" {
				return res
			}
		}

	} else {
		for i := len(runesOfStr) - 1; i > -1; i-- {
			res := helperForFindDigitInString(strIn, runesOfStr[i], i)
			if res != "" {
				return res
			}
		}
	}

	return ""
}

func helperForFindDigitInString(strIn string, runeIn rune, runeIndexInStr int) string {
	if unicode.IsDigit(runeIn) {
		return string(runeIn)
	}

	if slices.Contains(digitStartingChars, runeIn) {
		digCheck := CheckIfStringContainsDigitNameStartingAtPosition(strIn, runeIndexInStr)
		if digCheck != "" {
			return digCheck
		}
	}

	return ""
}

func CheckIfStringContainsDigitNameStartingAtPosition(strIn string, startingIndex int) string {
	for k, v := range digitToIntMap {
		if strings.Contains(strIn[startingIndex:], k); strings.Index(strIn[startingIndex:], k) == 0 {
			return v
		}
	}

	return ""
}
