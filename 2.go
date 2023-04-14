package main

import "fmt"

func main() {
	var x int
	fmt.Print("digite a quantidade de numeros: ")
	fmt.Scan(&x)
	lista := make([]int, x)

	for i := 0; i < x; i++ {
		fmt.Print("digite o elemento", i+1, ": ")
		fmt.Scan(&lista[i])
	}
	resultado, err := media(lista)
	if err != nil {
		fmt.Printf("nao existem numeros na lista")
	}
	fmt.Print(resultado)
}
func media(numeros []int) (int, error) {
	if len(numeros) == 0 {
		return 0, fmt.Errorf("nao existe numeros")
	}
	soma := 0
	for _, c := range numeros {
		soma += c
	}
	return soma / len(numeros), nil
}
