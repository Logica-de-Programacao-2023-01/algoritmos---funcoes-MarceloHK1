package marcelo

import "fmt"

func pares(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}
func oito(slice []int) (a []int, err error) {
	if len(slice) == 0 {
		return nil, fmt.Errorf("o slice esta vazio")
	}
	slicefinal := make([]int, 0)
	for _, i := range slice {
		if pares(i) == true {
			slicefinal = append(slicefinal, i)
		}
	}
	return slicefinal, nil

}
