package marcelo

import "fmt"

func junção(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", fmt.Errorf("o slice está vazio")
	}
	concatenada := ""
	for _, i := range slice {
		concatenada += i + ","
	}
	return concatenada, nil
}
