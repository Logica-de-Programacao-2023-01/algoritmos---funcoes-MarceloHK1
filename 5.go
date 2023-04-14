package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Print("digite o numero de elementos da sequencia: ")
	fmt.Scan(&x)

	resultado, err := Fibonacci(x)
	if err != nil {
		fmt.Print("a sequencia nao existe")
		return
	}
	fmt.Print(resultado)
}

func Fibonacci(n int) ([]int, error) {
	if n < 0 {
		return nil, fmt.Errorf("O número deve ser positivo")
	}
	seq := make([]int, n)
	if n >= 1 {
		seq[0] = 1
	}
	if n >= 2 {
		seq[1] = 1
	}
	for i := 2; i < n; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}
	return seq, nil
}
