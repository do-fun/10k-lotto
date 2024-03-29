package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	numberList := make([]int, 0)

	for len(numberList) < 6 {
		number := rand.Intn(45) + 1

		if !removeDuplicateNumber(numberList, number) {
			numberList = append(numberList, number)
		}
	}

	fmt.Println("Number List:", numberList)

}

func removeDuplicateNumber(targetNumberList []int, checkNumber int) bool {
	for _, targetNumber := range targetNumberList {
		if targetNumber == checkNumber {
			return true
		}
	}
	return false
}
