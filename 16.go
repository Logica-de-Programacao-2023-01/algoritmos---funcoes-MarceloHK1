package marcelo

import (
	"fmt"
	"strings"
)

func dezesseis(a string) (string, error) {
	if len(a) == 0 {
		return "", fmt.Errorf("a string é vazia")
	}
	vogais := "aeiouAEIOU"
	novastring := strings.ReplaceAll(a, vogais, "*")
	return novastring, nil
}
