package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func extractMultiply(mul string) int {
	reMul := regexp.MustCompile(`\d+`)
	numStrings := reMul.FindAllString(mul, -1)
	numInts := []int{}
	for _, numStr := range numStrings {
		numInt, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		numInts = append(numInts, numInt)
	}
	return numInts[0] * numInts[1]
}

// find the most recent instruction before this mul
func isMulEnabled(mulStartIdx int, dontStartIndexes []int, doStartIndexes []int) bool {
	lastInstruction := "do" // default state is enabled
	lastInstructionIdx := -1

	for _, dontIdx := range dontStartIndexes {
		if dontIdx < mulStartIdx && dontIdx > lastInstructionIdx {
			lastInstruction = "dont"
			lastInstructionIdx = dontIdx
		}
	}

	for _, doIdx := range doStartIndexes {
		if doIdx < mulStartIdx && doIdx > lastInstructionIdx {
			lastInstruction = "do"
			lastInstructionIdx = doIdx
		}
	}

	return lastInstruction == "do"
}

func main() {
	file, err := os.Open("day3/sample.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)

	fullText := ""
	for scanner.Scan() {
		fullText += scanner.Text()
	}

	muls := re.FindAllStringIndex(fullText, -1)
	dos := reDo.FindAllStringIndex(fullText, -1)
	donts := reDont.FindAllStringIndex(fullText, -1)

	dontStartIndexes := []int{}
	for _, dont := range donts {
		dontStartIndexes = append(dontStartIndexes, dont[0])
	}

	doStartIndexes := []int{}
	for _, do := range dos {
		doStartIndexes = append(doStartIndexes, do[0])
	}

	total := 0
	for _, mul := range muls {
		if isMulEnabled(mul[0], dontStartIndexes, doStartIndexes) {
			result := extractMultiply(fullText[mul[0]:mul[1]])
			total += result
		}
	}

	fmt.Println("Total:", total)
}
