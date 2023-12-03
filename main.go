package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/joshpaq/advent_of_code_2023/day1"
	"github.com/joshpaq/advent_of_code_2023/day2"
	"github.com/joshpaq/advent_of_code_2023/day3"
)

func main() {
	var day string
	var part string

	app := &cli.App{
		Name: "Advent of Code 2023",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "day, d",
				Usage:       "Advent of Code day to run",
				Destination: &day,
			},
			&cli.StringFlag{
				Name:        "part, p",
				Value:       "1",
				Usage:       "Advent of Code part of day to run",
				Destination: &part,
			},
		},
		Action: func(ctx *cli.Context) error {
			var result int

			switch day {
			case "1":
				switch part {
				case "1":
					result = day1.Part1()
				case "2":
					result = day1.Part2()
				}
			case "2":
				switch part {
				case "1":
					result = day2.Part1()
				case "2":
					result = day2.Part2()
				}
			case "3":
				switch part {
				case "1":
					result = day3.Part1()
				case "2":
					result = day3.Part2()
				}
			}

			fmt.Printf("Day %v Part %v: %v", day, part, result)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
