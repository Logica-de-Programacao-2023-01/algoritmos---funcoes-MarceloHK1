package marcelo

import "fmt"

func quinze(a int, b []int) (bool, error) {
	if len(b) == 0 {
		return false, fmt.Errorf("o slice é vazio")
	}
	for _, i := range b {
		if a == i {
			return true, nil
			break
		}
	}
	return false, nil
}
