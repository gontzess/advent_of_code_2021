package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Fish struct {
	DoB int
}

func (fish Fish) findChildren(nthDay int) []*Fish {
	var children []*Fish

	for day := fish.DoB + 8; day < nthDay; day += 7 {
		child := &Fish{DoB: day + 1}
		children = append(children, child)
	}

	for _, child := range children {
		children = append(children, child.findChildren(nthDay)...)
	}

	return children
}

func main() {
	// school := parseInput(sampleData())
	school := parseInput(steveData())

	// fmt.Println(school)
	// fmt.Println(countPopulation(school, 18))
	fmt.Println(countPopulation(school, 80))

}

func countPopulation(school []*Fish, nthDay int) int {
	count := len(school)
	for _, fish := range school {
		count += len(fish.findChildren(nthDay))
	}

	return count
}

func parseInput(inputFishSS string) []*Fish {
	fishSS := strings.Split(inputFishSS, ",")
	school := make([]*Fish, len(fishSS))

	for idx, fishSS := range fishSS {
		fishSSInt, _ := strconv.Atoi(fishSS)
		school[idx] = &Fish{fishSSInt - 8}
	}

	return school
}

func sampleData() string {
	return `3,4,3,1,2`
}

func steveData() string {
	return `1,1,3,1,3,2,1,3,1,1,3,1,1,2,1,3,1,1,3,5,1,1,1,3,1,2,1,1,1,1,4,4,1,2,1,2,1,1,1,5,3,2,1,5,2,5,3,3,2,2,5,4,1,1,4,4,1,1,1,1,1,1,5,1,2,4,3,2,2,2,2,1,4,1,1,5,1,3,4,4,1,1,3,3,5,5,3,1,3,3,3,1,4,2,2,1,3,4,1,4,3,3,2,3,1,1,1,5,3,1,4,2,2,3,1,3,1,2,3,3,1,4,2,2,4,1,3,1,1,1,1,1,2,1,3,3,1,2,1,1,3,4,1,1,1,1,5,1,1,5,1,1,1,4,1,5,3,1,1,3,2,1,1,3,1,1,1,5,4,3,3,5,1,3,4,3,3,1,4,4,1,2,1,1,2,1,1,1,2,1,1,1,1,1,5,1,1,2,1,5,2,1,1,2,3,2,3,1,3,1,1,1,5,1,1,2,1,1,1,1,3,4,5,3,1,4,1,1,4,1,4,1,1,1,4,5,1,1,1,4,1,3,2,2,1,1,2,3,1,4,3,5,1,5,1,1,4,5,5,1,1,3,3,1,1,1,1,5,5,3,3,2,4,1,1,1,1,1,5,1,1,2,5,5,4,2,4,4,1,1,3,3,1,5,1,1,1,1,1,1`
}
