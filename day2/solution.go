package day2

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input
var input string

func splitCube(cube string) (int, string) {
	split := strings.Split(strings.TrimSpace(cube), " ")
	count, _ := strconv.Atoi(split[0])
	return count, split[1]
}

func sum(arr []int) int {
	var sum = 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

const RED_MAX = 12
const GREEN_MAX = 13
const BLUE_MAX = 14

func Part1() int {
	var red = 0
	var green = 0
	var blue = 0
	var possible []int

	for i, line := range strings.Split(input, "\r\n") {
		game := strings.Split(line, ":")[1]

		for _, pull := range strings.Split(game, ";") {
			for _, cube := range strings.Split(pull, ",") {
				count, color := splitCube(cube)
				switch color {
				case "blue":
					if count > blue {
						blue = count
					}
				case "green":
					if count > green {
						green = count
					}
				case "red":
					if count > red {
						red = count
					}
				}
			}
		}

		if BLUE_MAX >= blue && GREEN_MAX >= green && RED_MAX >= red {
			possible = append(possible, i+1)
		}

		blue = 0
		green = 0
		red = 0
	}

	return sum(possible)
}

func Part2() int {
	var red = 0
	var green = 0
	var blue = 0
	var powers []int

	for _, line := range strings.Split(input, "\r\n") {
		game := strings.Split(line, ":")[1]

		for _, pull := range strings.Split(game, ";") {
			for _, cube := range strings.Split(pull, ",") {
				count, color := splitCube(cube)
				switch color {
				case "blue":
					if count > blue {
						blue = count
					}
				case "green":
					if count > green {
						green = count
					}
				case "red":
					if count > red {
						red = count
					}
				}
			}
		}

		power := blue * red * green
		powers = append(powers, power)

		blue = 0
		green = 0
		red = 0
	}

	return sum(powers)
}
