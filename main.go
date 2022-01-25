package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 기본 로또번호 6개 생성
func setLottoNumberList() []int {
	numberList := []int{}
	rand.Seed(int64(time.Now().UnixNano()))
	for i := 0; i < 6; i++ {
		randNum := rand.Intn(45) + 1

		numberList = append(numberList, randNum)
	}

	return numberList
}

//중복제거 후 빈자리에 랜덤숫자 다시 생성
func removeDuplicateNumber(numberList []int) []int {
	allKeys := make(map[int]bool)
	removedNumberList := []int{}

	for _, item := range numberList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			removedNumberList = append(removedNumberList, item)
		} else {
			rand.Seed(int64(time.Now().UnixNano()))
			newRandNum := rand.Intn(45) + 1
			removedNumberList = append(removedNumberList, newRandNum)
		}
	}

	return removedNumberList
}

func main() {
	initNumberList := setLottoNumberList()
	fmt.Println("initNumberList : ", initNumberList)

	finalNumberList := removeDuplicateNumber(initNumberList)
	fmt.Println("finalNumberList : ", finalNumberList)

}
