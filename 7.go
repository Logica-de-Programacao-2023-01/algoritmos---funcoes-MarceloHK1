package marcelo

import (
	"fmt"
)

func somar(a int) int {
	return a + 2
}
func sete(a []int) (b []int, err error) {
	if len(a) == 0 {
		return nil, fmt.Errorf("o slice esta vazio")
	}
	slicefinal := make([]int, len(a))
	for _, i := range a {
		resultados := somar(i)
		slicefinal = append(slicefinal, resultados)
	}
	return slicefinal, nil
}
