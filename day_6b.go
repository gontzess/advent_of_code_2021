package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Fish struct {
	DoB           int
	childrenCount int64
}

func (fish Fish) countChildren(nthDay int) int64 {
	for day := fish.DoB + 8; day < nthDay; day += 7 {
		child := &Fish{DoB: day + 1, childrenCount: 0}
		fish.childrenCount += 1
		fish.childrenCount += child.countChildren(nthDay)
	}

	return fish.childrenCount
}

func (fish Fish) offset() int {
	return fish.DoB + 8
}

func main() {
	// school := parseInput(sampleData())
	// school := parseInput(testData())
	school := parseInput(steveData())

	// fmt.Println(school)
	// fmt.Println(countPopulation(school, 18))
	// fmt.Println(countPopulation(school, 80))
	fmt.Println(countPopulation(school, 256))
}

func countPopulation(school []*Fish, nthDay int) int64 {
	nthTemp := make([]int64, 9)
	fmt.Println(nthTemp)
	count := int64(len(school))
	for _, fish := range school {
		if nthTemp[fish.offset()] == int64(0) {
			nthTemp[fish.offset()] = fish.countChildren(nthDay)
		}
		count += nthTemp[fish.offset()]
	}

	fmt.Println(nthTemp)
	return count
}

func parseInput(inputFishSS string) []*Fish {
	fishSS := strings.Split(inputFishSS, ",")
	school := make([]*Fish, len(fishSS))

	for idx, fishSS := range fishSS {
		fishSSInt, _ := strconv.Atoi(fishSS)
		school[idx] = &Fish{DoB: fishSSInt - 8, childrenCount: 0}
	}

	return school
}

func sampleData() string {
	return `3,4,3,1,2`
}

func testData() string {
	return `5`
}

func steveData() string {
	return `1,1,3,1,3,2,1,3,1,1,3,1,1,2,1,3,1,1,3,5,1,1,1,3,1,2,1,1,1,1,4,4,1,2,1,2,1,1,1,5,3,2,1,5,2,5,3,3,2,2,5,4,1,1,4,4,1,1,1,1,1,1,5,1,2,4,3,2,2,2,2,1,4,1,1,5,1,3,4,4,1,1,3,3,5,5,3,1,3,3,3,1,4,2,2,1,3,4,1,4,3,3,2,3,1,1,1,5,3,1,4,2,2,3,1,3,1,2,3,3,1,4,2,2,4,1,3,1,1,1,1,1,2,1,3,3,1,2,1,1,3,4,1,1,1,1,5,1,1,5,1,1,1,4,1,5,3,1,1,3,2,1,1,3,1,1,1,5,4,3,3,5,1,3,4,3,3,1,4,4,1,2,1,1,2,1,1,1,2,1,1,1,1,1,5,1,1,2,1,5,2,1,1,2,3,2,3,1,3,1,1,1,5,1,1,2,1,1,1,1,3,4,5,3,1,4,1,1,4,1,4,1,1,1,4,5,1,1,1,4,1,3,2,2,1,1,2,3,1,4,3,5,1,5,1,1,4,5,5,1,1,3,3,1,1,1,1,5,5,3,3,2,4,1,1,1,1,1,5,1,1,2,5,5,4,2,4,4,1,1,3,3,1,5,1,1,1,1,1,1`
}
