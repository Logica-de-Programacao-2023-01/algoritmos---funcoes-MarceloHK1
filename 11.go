package marcelo

import "fmt"

func onze(a int) (bool, error) {
	if a < 0 {
		return false, fmt.Errorf("o numero é negativo")
	}
	if a < 2 {
		return false, nil
	}
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false, nil
		}
	}
	return true, nil

}
