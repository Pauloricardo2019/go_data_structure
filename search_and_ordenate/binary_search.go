package search_and_ordenate

import (
	"fmt"
	"math/rand"
	"time"
)

func binary_search(numbers []int, toSearch int) {
	finish := len(numbers) - 1
	inicio := 0

	for inicio <= finish {
		middle := (finish + inicio) / 2
		actually := numbers[middle]

		if actually == toSearch {
			fmt.Println("ENCONTROU!!!")
			fmt.Println(toSearch)
			break
		}

		if actually > toSearch {
			fmt.Println("O numero é maior!!!")
			finish = middle - 1
		}

		if actually < toSearch {
			fmt.Println("O numero é menor!!!")
			inicio = middle + 1
		}

	}

}

func generateRandomNumbers() []int {
	rand.Seed(time.Now().UnixNano())

	randomNumbers := make([]int, 0)

	numCount := 100000
	maxNum := 9865498987

	for i := 0; i < numCount; i++ {
		randomNumber := rand.Intn(maxNum)
		randomNumbers = append(randomNumbers, randomNumber)
	}
	return randomNumbers
}

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {

	numbers := generateRandomNumbers()
	numbers = bubbleSort(numbers)
	randIdx := rand.Intn(len(numbers))
	toSearch := numbers[randIdx]
	fmt.Println("Number to search: ", toSearch)

	initTime := time.Now()
	binary_search(numbers, toSearch)

	fmt.Println(fmt.Sprintf("Time total: %v", time.Now().Sub(initTime)))

}
