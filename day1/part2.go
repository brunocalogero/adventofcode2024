package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("day1/sample.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	listOne := []int{}
	listTwo := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "   ")

		itemInt1, err := strconv.Atoi(items[0])
		if err != nil {
			panic(err)
		}

		itemInt2, err := strconv.Atoi(items[1])
		if err != nil {
			panic(err)
		}

		listOne = append(listOne, itemInt1)
		listTwo = append(listTwo, itemInt2)
	}

	// Fill up the frequency map
	counter := map[int]int{}
	for _, elem2 := range listTwo {
		if _, ok := counter[elem2]; !ok {
			counter[elem2] = 1
		} else {
			counter[elem2]++
		}
	}

	// Iterate through first list to calculate similarity score
	similarityScore := 0
	for _, elem1 := range listOne {
		if _, ok := counter[elem1]; ok {
			similarityScore = similarityScore + elem1*counter[elem1]
		}
	}

	fmt.Println(similarityScore)
}
