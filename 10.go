package marcelo

import "fmt"

func dez(a []int) (b []int, err error) {
	if len(a) == 0 {
		return nil, fmt.Errorf("o slice é vazio")
	}
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a, nil
}
