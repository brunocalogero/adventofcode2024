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

func isValidSequence(nums []int, increasing bool) bool {
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		if increasing {
			if diff < 1 || diff > 3 {
				return false
			}
		} else {
			if diff > -1 || diff < -3 {
				return false
			}
		}
	}
	return true
}

func isSafeWithDampener(nums []int) bool {
	// check safe without removing anything
	if isValidSequence(nums, true) || isValidSequence(nums, false) {
		return true
	}

	// check safe by removing single
	for i := 0; i < len(nums); i++ {
		withoutI := append([]int{}, nums[:i]...)
		withoutI = append(withoutI, nums[i+1:]...)

		if isValidSequence(withoutI, true) || isValidSequence(withoutI, false) {
			return true
		}
	}

	return false
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
		if isSafeWithDampener(lineItems) {
			safeCounter++
		}
	}

	fmt.Println(safeCounter)
}
