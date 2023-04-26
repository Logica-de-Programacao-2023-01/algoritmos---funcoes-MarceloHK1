package marcelo

import "fmt"

func vinte(a []string) ([]string, error) {
	if len(a) == 0 {
		return nil, fmt.Errorf("o slice está vazio")
	}
	novoslice := make([]string, 0)
	for _, i := range a {
		if len(i) > 5 {
			novoslice = append(novoslice, i)
		}
	}
	return novoslice, nil
}
