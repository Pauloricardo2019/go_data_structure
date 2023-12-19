package linked_list

import (
	"fmt"
)

type node struct {
	value int
	prox  *node
}

func insertOnInit(lista **node, value int) {
	novo := &node{}

	if novo != nil {
		novo.value = value
		novo.prox = *lista
		*lista = novo

	} else {
		fmt.Println("Erro ao criar novo nó!")
	}
}

func insertOnFinish(lista **node, value int) {
	novo := &node{}
	aux := &node{}

	if novo != nil {
		novo.value = value
		if *lista == nil {
			novo.prox = *lista
			*lista = novo
		} else {
			aux = *lista
			//lista: 5 6
			//inserir o 7
			for aux.prox != nil {
				aux = aux.prox
			}
			//aux == 6
			aux.prox = novo
			novo.prox = nil
		}

	} else {
		fmt.Println("Erro ao criar novo nó!")
	}
}

func insertOnMiddle(list **node, value, ant int) {
	novo := &node{}
	aux := &node{}

	if novo != nil {
		novo.value = value
		if *list == nil {
			novo.prox = *list
			*list = novo
		} else {
			aux = *list
			// lista => 1 , 2 , 3 , 4
			// ant = 2
			//value = 5
			for aux.value != ant && aux.prox != nil {
				aux = aux.prox
			}
			novo.prox = aux.prox
			aux.prox = novo
		}

	} else {
		fmt.Println("Erro ao criar novo nó!")
	}

}

func printList(list *node) {
	fmt.Println("Imprimindo lista!")
	for list != nil {
		fmt.Printf("%d -> ", list.value)
		list = list.prox
	}
	fmt.Println("\nAcabou! '-'")
}

func main() {

	var lista *node

	//insertOnInit(&lista, 1)
	//insertOnInit(&lista, 2)
	//insertOnInit(&lista, 3)

	insertOnFinish(&lista, 1)
	insertOnFinish(&lista, 2)
	insertOnFinish(&lista, 3)
	insertOnMiddle(&lista, 5, 2)
	insertOnInit(&lista, 9)

	printList(lista)

}
