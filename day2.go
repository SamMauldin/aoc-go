package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func day2part1(commands []command) int {
	horizontal_pos := 0
	depth := 0

	for _, v := range commands {
		if v.direction == "forward" {
			horizontal_pos += v.magnitude
		} else if v.direction == "down" {
			depth += v.magnitude
		} else if v.direction == "up" {
			depth -= v.magnitude
		} else {
			log.Fatalf("Unknown direction %s", v.direction)
		}
	}

	return horizontal_pos * depth
}

func day2part2(commands []command) int {
	horizontal_pos := 0
	depth := 0
	aim := 0

	for _, v := range commands {
		if v.direction == "forward" {
			horizontal_pos += v.magnitude
			depth += aim * v.magnitude
		} else if v.direction == "down" {
			aim += v.magnitude
		} else if v.direction == "up" {
			aim -= v.magnitude
		} else {
			log.Fatalf("Unknown direction %s", v.direction)
		}
	}

	return horizontal_pos * depth
}

type command struct {
	direction string
	magnitude int
}

func day2() {
	content, err := ioutil.ReadFile(("day2.txt"))

	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	command_strs := strings.Split(strings.TrimSpace(input), "\n")

	commands := make([]command, len(command_strs))
	for i, command_str := range command_strs {
		parts := strings.Split(command_str, " ")

		magnitude, err := strconv.Atoi(parts[1])

		if err != nil {
			log.Fatal((err))
		}

		commands[i] = command{
			direction: parts[0],
			magnitude: magnitude,
		}
	}

	fmt.Println("--- Day 2 ---")
	fmt.Println("-- Part 1 --")
	fmt.Println(day2part1(commands))
	fmt.Println("-- Part 2 --")
	fmt.Println(day2part2(commands))
}
