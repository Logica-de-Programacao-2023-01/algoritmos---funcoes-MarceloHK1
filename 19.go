package marcelo

import "fmt"

func dezenove(a int) ([]int, error) {
	if a < 0 {
		return nil, fmt.Errorf("o número é negativo")
	}
	primos := make([]int, 0)
	for i := 2; i <= a; i++ {
		primo := true
		for j := 2; j*j <= a; j++ {
			if i%j == 0 {
				primo = false
				break
			}
		}
		if primo {
			primos = append(primos, i)
		}
	}
	return primos, nil
}
