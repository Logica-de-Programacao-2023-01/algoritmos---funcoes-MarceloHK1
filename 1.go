package main

import "fmt"

func main() {
	var x, y float64
	fmt.Print("digite o primeiro numero: ")
	fmt.Scan(&x)
	fmt.Print("digite o segundo numero: ")
	fmt.Scan(&y)

	resultado, err := divide(x, y)
	if err != nil {
		fmt.Printf("Ocorreu um erro ao dividir x e y: %s\n", err)
	}
	fmt.Println(resultado)
}
func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("nao se divide um numero por 0")
	}
	return x / y, nil
}
