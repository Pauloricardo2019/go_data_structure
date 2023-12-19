package linked_list

import (
	"fmt"
	"math/rand"
	"time"
)

type nodeDuplicate struct {
	value int
	prox  *nodeDuplicate
	ant   *nodeDuplicate
}

func insert_on_init(list **nodeDuplicate, value int) {
	newNode := &nodeDuplicate{}

	if newNode != nil {
		newNode.value = value
		newNode.ant = nil
		newNode.prox = *list
		*list = newNode

	} else {
		fmt.Printf("error on allocate var in heap memory.")
	}
}

func insert_on_finish(list **nodeDuplicate, value int) {
	newNode := &nodeDuplicate{}
	aux := &nodeDuplicate{}

	if newNode != nil {
		newNode.value = value
		newNode.prox = nil
		if *list == nil {
			newNode.ant = nil
			*list = newNode
		} else {
			aux = *list
			for aux.prox != nil {
				aux = aux.prox
			}
			newNode.ant = aux
			aux.prox = newNode
		}

	} else {
		fmt.Printf("error on allocate var in heap memory.")
	}
}

func insert_on_middle(list **nodeDuplicate, ant, value int) {
	newNode := &nodeDuplicate{}
	aux := &nodeDuplicate{}

	if newNode != nil {
		newNode.value = value
		if *list == nil {
			newNode.ant = nil
			newNode.prox = *list
			*list = newNode
		}
		aux = *list
		// 9 5 7
		// inserir o 12 depois do 5
		for aux.prox != nil && aux.value != ant {
			aux = aux.prox
		}
		newNode.ant = aux
		newNode.prox = aux.prox
		aux.prox = newNode

	} else {
		fmt.Printf("error on allocate var in heap memory.")
	}

}

func insert_on_ordenate(list **nodeDuplicate, value int) {
	newNode := &nodeDuplicate{value: value}
	aux := *list

	if newNode != nil {
		if aux == nil || aux.value >= value {
			newNode.prox = aux
			newNode.ant = nil
			if aux != nil {
				aux.ant = newNode
			}
			*list = newNode
		} else {
			for aux.prox != nil && aux.prox.value < value {
				aux = aux.prox
			}
			newNode.prox = aux.prox
			newNode.ant = aux
			if aux.prox != nil {
				aux.prox.ant = newNode
			}
			aux.prox = newNode
		}
	} else {
		fmt.Printf("Error allocating memory on heap.")
	}
}

func printDuplicateList(list *nodeDuplicate) {
	fmt.Printf("Print list\n")

	for list != nil {
		fmt.Printf("%d -> ", list.value)
		list = list.prox
	}
	fmt.Printf("\nFinish print...")

}

//func main() {
//	timeInit := time.Now()
//	var list *nodeDuplicate
//	numbers := generateRandomNumbers()
//	for _, number := range numbers {
//		insert_on_ordenate(&list, number)
//	}
//	//printList(list)
//	//n := bubbleSort(numbers)
//	//fmt.Println(n)
//	fmt.Println(fmt.Sprintf("time total: %v", time.Now().Sub(timeInit)))
//
//}

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Troca os elementos se estiverem fora de ordem
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
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
