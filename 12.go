package marcelo

import (
	"fmt"
	"strings"
)

func doze(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", fmt.Errorf("o slice esta vazio")
	}
	maiusculas := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	stringfinal := ""
	for i := 0; i < len(slice); i++ {
		if strings.Contains(slice[i], maiusculas) == true {
			stringfinal += slice[i]
		}
	}
	return stringfinal, nil

}
