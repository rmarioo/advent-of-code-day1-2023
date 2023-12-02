package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"strconv"
)

func Trebuchet(supplier LinesSupplier) int {

	lines := supplier.Lines()

	/* imperative solution */
	/*
		sum := 0
			for _, line := range lines {
				sum += DigitsIn(line)
			}
	*/

	/* functional solution 1: map lines to digits and anfter sum */
	/*	var digits = lo.Map(lines, func(line string, _ int) int { return DigitsIn(line) })
		sum := lo.Sum(digits)*/

	/* functional solution 2 : reduce lines to digits and sum */
	sum := lo.Reduce(lines, func(agg int, line string, _ int) int { return agg + DigitsIn(line) }, 0)

	return sum

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
func DigitsIn(entry string) int {

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

	sprintf := fmt.Sprintf("%d%d", firstDigit, lastDigit)
	i, _ := strconv.Atoi(sprintf)
	return i
}
