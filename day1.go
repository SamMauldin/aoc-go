package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func part1(depths []int) int {
	increases := 0

	for i, v := range depths {
		if i != 0 && v > depths[i-1] {
			increases++
		}
	}

	return increases
}

func part2(depths []int) int {
	increases := 0

	for i, v := range depths {
		if i >= 3 && v > depths[i-3] {
			increases++
		}
	}

	return increases
}

func main() {
	content, err := ioutil.ReadFile(("day1.txt"))

	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	depths_strs := strings.Split(strings.TrimSpace(input), "\n")

	depths := make([]int, len(depths_strs))
	for i, depth_str := range depths_strs {
		depth, err := strconv.Atoi(depth_str)

		if err != nil {
			log.Fatal((err))
		}

		depths[i] = depth
	}

	fmt.Println("-- Part 1 --")
	fmt.Println(part1(depths))
	fmt.Println("-- Part 2 --")
	fmt.Println(part2(depths))
}
