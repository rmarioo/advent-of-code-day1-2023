package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Trebuchet(supplier LinesSupplier) int {

	lines := supplier.Lines()

	sum := lo.Reduce(lines, func(acc int, line string, _ int) int { return acc + DigitsIn(line) }, 0)

	return sum

	/* imperative solution */

	/*	sum := 0
		for _, line := range lines {

			var in = DigitsIn(line)
			//fmt.Println("calculated %d", in)
			sum += in
		}*/

	/* functional solution 1: map lines to digits and after sum */
	/*
		var digits []int = lo.Map(lines, func(line string, _ int) int { return DigitsIn(line) })
		sum := lo.Sum(digits)
	*/
	/* functional solution 2 : reduce lines to digits and sum */

}

type LinesSupplier interface {
	Lines() []string
}

type FromFileLinesSupplier struct {
	fileName string
}

func (f FromFileLinesSupplier) Lines() []string {
	return ReadFileLines(f.fileName)
}
func ReadFileLines(fileName string) []string {

	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}

func main() {

	trebuchet := Trebuchet(FromFileLinesSupplier{fileName: "input.txt"})

	fmt.Println("solution ", trebuchet)
}

func DigitsIn(entryInput string) int {

	stringWithOutWords := replaceWordNumbers(strings.Clone(entryInput))

	firstDigit, lastDigit := retrieveFirstAndLastDigit(stringWithOutWords)

	return concatDigits(firstDigit, lastDigit)
}

func concatDigits(firstDigit int, lastDigit int) int {
	sprintf := fmt.Sprintf("%d%d", firstDigit, lastDigit)
	i, _ := strconv.Atoi(sprintf)
	return i
}

func retrieveFirstAndLastDigit(entry string) (int, int) {
	var firstDigit = -1
	var lastDigit = -1
	for _, ch := range entry {
		n, err := strconv.Atoi(string(ch))
		if err == nil {
			if firstDigit == -1 {
				firstDigit = n
			}
			lastDigit = n
		}
	}
	return firstDigit, lastDigit
}

func replaceWordNumbers(entry string) string {
	numberWordReplacement := findOrderedNumberWordReplacement(entry)
	for _, nw := range numberWordReplacement {
		entry = strings.ReplaceAll(entry, nw.numberWord.numberWord, fmt.Sprint(nw.numberWord.value))
	}
	return entry
}

var numberWords = []NumberWord{
	{"nineight", 98},
	{"eighthree", 83},
	{"eightwo", 82},
	{"sevenine", 79},
	{"fiveight", 58},
	{"threeight", 38},
	{"twone", 21},
	{"oneight", 18},

	{"nine", 9},
	{"eight", 8},
	{"seven", 7},
	{"six", 6},
	{"five", 5},
	{"four", 4},
	{"three", 3},
	{"two", 2},
	{"one", 1},
}

type NumberWord struct {
	numberWord string
	value      int
}

type NumberWordReplacement struct {
	numberWord NumberWord
	position   int
}

func findOrderedNumberWordReplacement(s string) []NumberWordReplacement {
	var numberWordsIdx []NumberWordReplacement

	for _, nw := range numberWords {
		idx := strings.Index(s, nw.numberWord)
		if idx != -1 {
			numberWordsIdx = append(numberWordsIdx, NumberWordReplacement{numberWord: nw, position: idx})
		}
	}

	sort.Slice(numberWordsIdx, func(i, j int) bool {
		return numberWordsIdx[i].position < numberWordsIdx[j].position
	})
	return numberWordsIdx
}
