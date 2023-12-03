package day3

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input
var input string

func isSymbol(character string) bool {
	return character != "." && character != "1" && character != "2" && character != "3" && character != "4" && character != "5" && character != "6" && character != "7" && character != "8" && character != "9" && character != "0"
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type Coordinate struct {
	X int
	Y int
}

func isXAdjacent(symbolX, numberXStart, numberXEnd int) bool {
	return symbolX >= numberXStart-1 && symbolX <= numberXEnd+1
}

func convert(value string) int {
	number, _ := strconv.Atoi(value)
	return number
}

func Part1() int {
	total := 0

	var symbols []Coordinate

	lines := strings.Split(input, "\r\n")

	for y, line := range lines {
		for x, character := range strings.Split(line, "") {
			if isSymbol(character) {
				symbols = append(symbols, Coordinate{X: x, Y: y})
			}
		}
	}

	numberMatcher := regexp.MustCompile("\\d+")
	for _, symbol := range symbols {
		top := max(symbol.Y-1, 0)
		bottom := min(symbol.Y+1, len(lines))
		for _, line := range lines[top : bottom+1] {
			for _, number := range numberMatcher.FindAllIndex([]byte(line), -1) {
				if isXAdjacent(symbol.X, number[0], number[1]-1) {
					total += convert(line[number[0]:number[1]])
				}
			}
		}
	}

	return total
}

func Part2() int {
	total := 0
	var gears []Coordinate

	lines := strings.Split(input, "\r\n")

	for y, line := range lines {
		for x, character := range strings.Split(line, "") {
			if character == "*" {
				gears = append(gears, Coordinate{X: x, Y: y})
			}
		}
	}

	numberMatcher := regexp.MustCompile("\\d+")
	for _, gear := range gears {
		top := max(gear.Y-1, 0)
		bottom := min(gear.Y+1, len(lines))
		var adjacentGears []int
		for _, line := range lines[top : bottom+1] {
			for _, number := range numberMatcher.FindAllIndex([]byte(line), -1) {
				if isXAdjacent(gear.X, number[0], number[1]-1) {
					adjacentGears = append(adjacentGears, convert(line[number[0]:number[1]]))
				}
			}
		}
		if len(adjacentGears) == 2 {
			total += adjacentGears[0] * adjacentGears[1]
		}
	}

	return total
}
