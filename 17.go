package marcelo

import (
	"fmt"
)

func dezessete(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", fmt.Errorf("o slice é vazio")
	}
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	stringfinal := ""
	for _, x := range slice {
		stringfinal += x
	}
	return stringfinal, nil
}
