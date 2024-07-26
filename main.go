package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	tempNumberList := make([]int, 0)

	for i := 0; i < 5; i++ {
		numberList := make([]int, 0)
		for len(numberList) < 6 {
			number := rand.Intn(45) + 1

			if !removeDuplicateNumber(numberList, number) {
				numberList = append(numberList, number)
			}
		}

		// 오름차순 정렬
		slices.Sort(numberList)
		fmt.Println("Number List:", numberList)

		tempNumberList = append(tempNumberList, numberList...)

	}

	uniqueTempNumberList := removeDuplicates(tempNumberList)
	// fmt.Println("Unique Temp Number List:", uniqueTempNumberList)

	for i := 0; i < 5; i++ {
		newNumberList := make([]int, 0)
		for len(newNumberList) < 6 {
			newNumber := rand.Intn(45) + 1

			if !removeDuplicateNumber(uniqueTempNumberList, newNumber) && !removeDuplicateNumber(newNumberList, newNumber) {
				newNumberList = append(newNumberList, newNumber)
			}
		}

		slices.Sort(newNumberList)
		fmt.Println("New Number List:", newNumberList)
	}

}

func removeDuplicateNumber(targetNumberList []int, checkNumber int) bool {
	for _, targetNumber := range targetNumberList {
		if targetNumber == checkNumber {
			return true
		}
	}
	return false
}

func removeDuplicates(numbers []int) []int {
	seen := make(map[int]bool)
	uniqueNumbers := make([]int, 0)

	for _, number := range numbers {
		if !seen[number] {
			seen[number] = true
			uniqueNumbers = append(uniqueNumbers, number)
		}
	}

	return uniqueNumbers
}
