package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("digite uma string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	resultado, err := frase(str)
	if err != nil {
		fmt.Print("string inexistente")
		return
	}
	fmt.Print(resultado)
}

func frase(x string) (int, error) {
	if len(x) == 0 {
		return 0, fmt.Errorf("string inexistente")
	}
	return len(x), nil
}
