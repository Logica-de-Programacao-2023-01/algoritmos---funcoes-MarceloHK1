package marcelo

import "fmt"

func dezoitoa(a int) int {
	return a + 2
}
func dezoito(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, fmt.Errorf("o slice é vazio")
	}
	novoslice := make([]int, 0)
	for _, i := range slice {
		resultado := dezoitoa(i)
		novoslice = append(novoslice, resultado)
	}
	soma := 0
	for _, j := range novoslice {
		soma += j
	}
	return soma, nil
}
