package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// positions := parseInput(readData("data/day_7_sampledata.txt"))
	positions := parseInput(readData("data/day_7_stevedata.txt"))
	sort.Ints(positions)
	// fmt.Println(positions)

	bestCost, bestPosition := findLowestCost(tally(positions), filter(positions))
	fmt.Println(bestCost, bestPosition)

	fmt.Println(median(positions))
}

func filter(sortedPositions []int) []int {
	var filteredPositions []int
	var lastPosition int
	for idx, position := range sortedPositions {
		if idx != 0 && position == lastPosition {
			continue
		}
		filteredPositions = append(filteredPositions, position)
		lastPosition = position
	}

	return filteredPositions
}

func findLowestCost(tally map[int]int, sortedPositions []int) (int, int) {
	bestCost := 10000000000
	bestPosition := -1
	start := sortedPositions[0]
	end := sortedPositions[len(sortedPositions)-1]
	for position := start; position <= end; position++ {
		cost := calcCost(tally, position)
		if cost < bestCost {
			bestCost, bestPosition = cost, position
			// fmt.Println(bestCost, position)
		}
	}

	return bestCost, bestPosition
}

func calcCost(tally map[int]int, target int) int {
	cost := 0

	for position, count := range tally {
		cost += int(math.Abs(float64(target-position))) * count
	}

	return cost
}

func median(sortedPositions []int) int {
	idx := len(sortedPositions) / 2
	return sortedPositions[idx]
}

func tally(sortedPositions []int) map[int]int {
	positionsTally := make(map[int]int)
	for _, position := range sortedPositions {
		positionsTally[position] += 1
	}

	return positionsTally
}

func parseInput(inputStr string) []int {
	inputStrs := strings.Split(inputStr, ",")
	positions := make([]int, len(inputStrs))

	for idx, position := range inputStrs {
		positionInt, _ := strconv.Atoi(position)
		positions[idx] = positionInt
	}

	return positions
}

func readData(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(data))
}
