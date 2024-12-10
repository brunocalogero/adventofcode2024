package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func strSliceToIntSlice(strSlice []string) []int {
	intSlice := []int{}
	for _, elem := range strSlice {
		intElem, err := strconv.Atoi(elem)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, intElem)
	}

	return intSlice
}

func isSafeAdj(a int, b int) bool {
	if a > b {
		diff := a - b
		if diff >= 1 && diff <= 3 {
			return true
		}
	}
	if a < b {
		diff := b - a
		if diff >= 1 && diff <= 3 {
			return true
		}
	}

	return false
}

func isDescreasing(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		if !(report[i] > report[i+1] && isSafeAdj(report[i], report[i+1])) {
			return false
		}
	}
	return true
}

func isIncreasing(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		if !(report[i] < report[i+1] && isSafeAdj(report[i], report[i+1])) {
			return false
		}
	}

	return true
}

func main() {

	file, err := os.Open("day2/sample.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	safeCounter := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineItems := strSliceToIntSlice(strings.Split(line, " "))

		if isDescreasing(lineItems) || isIncreasing(lineItems) {
			safeCounter++
		}
	}

	fmt.Println(safeCounter)
}
