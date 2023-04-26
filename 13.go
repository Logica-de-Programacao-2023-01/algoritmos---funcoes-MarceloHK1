package marcelo

import (
	"fmt"
	"strconv"
)

func treze(a int) (int, error) {
	if a < 0 {
		return 0, fmt.Errorf("o numero é negativo")
	}
	strnumeros := strconv.Itoa(a)
	soma := 0
	for _, i := range strnumeros {
		digito, err := strconv.Atoi(string(i))
		if err != nil {
			return 0, err
		}
		soma += digito
	}
	return soma, nil

}
