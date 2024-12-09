package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	sort.Ints(listOne)
	sort.Ints(listTwo)

	distance := 0
	for i := range listOne {
		if listOne[i] > listTwo[i] {
			distance = distance + listOne[i] - listTwo[i]
		} else {
			distance = distance + listTwo[i] - listOne[i]
		}
	}

	fmt.Println(distance)
}
