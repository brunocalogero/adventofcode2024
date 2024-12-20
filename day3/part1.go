package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("day3/sample.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	// NOTE: looks like input can be scanned by line.
	totalVal := 0
	for scanner.Scan() {
		line := scanner.Text()

		// find all the muls
		// NOTE: could make it better to only consider 3 digit per number
		re := regexp.MustCompile(`mul\(\d+,\d+\)`)
		lineMuls := re.FindAllString(line, -1)

		// for each mul in line extract numbers, convert, multiply & add to total
		reMul := regexp.MustCompile(`\d+`)
		for _, mul := range lineMuls {
			// extract
			numStrings := reMul.FindAllString(mul, -1)
			// convert
			numInts := []int{}
			for _, numStr := range numStrings {
				numInt, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				numInts = append(numInts, numInt)
			}
			// multiply
			totalMul := numInts[0] * numInts[1]
			// add to total
			totalVal = totalVal + totalMul
		}
	}

	fmt.Println(totalVal)
}
