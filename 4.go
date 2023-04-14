package main

import "fmt"

func main() {
	var size int
	fmt.Print("digite o tamanho do slice: ")
	fmt.Scan(&size)
	slice := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Print("digite o elemento", i+1, ": ")
		fmt.Scan(&slice[i])
	}
	resultado, err := list(slice)
	if err != nil {
		fmt.Print("o slice é vazio")
		return
	}
	fmt.Print(resultado)
}

func list(lista []int) (int, error) {
	if len(lista) == 0 {
		return 0, fmt.Errorf("o slice é vazio")
	}
	menor := lista[0]
	for _, c := range lista {
		if c < menor {
			menor = c
		}
	}
	return menor, nil
}
