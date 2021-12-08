package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// school := parseInput(readData("data/day_6_sampledata.txt"))
	school := parseInput(readData("data/day_6_stevedata.txt"))
	fmt.Println("starting school:", school)

	grow(school, 80)
	fmt.Println("ending school:", school)
	count := countPopulation(school)
	fmt.Println("school count:", count)
}

func countPopulation(school *[]int) int {
	sum := 0
	for _, fishCount := range *school {
		sum += fishCount
	}

	return sum
}

func grow(school *[]int, nthDay int) *[]int {
	if nthDay > 0 {
		tempZeros := (*school)[0]
		for fishState := 1; fishState <= 8; fishState++ {
			(*school)[fishState-1] = (*school)[fishState]
		}
		(*school)[6] += tempZeros
		(*school)[8] = tempZeros

		return grow(school, nthDay-1)
	}

	return school
}

func parseInput(inputFishSS string) *[]int {
	fishSS := strings.Split(inputFishSS, ",")
	school := make([]int, 9)

	for _, fishSS := range fishSS {
		fishSSInt, _ := strconv.Atoi(fishSS)
		school[fishSSInt] += 1
	}

	return &school
}

func readData(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(data))
}
