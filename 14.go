package marcelo

import "fmt"

func catorze(a []int, b []int) (c []int, err error) {
	if len(a) == 0 || len(b) == 0 {
		return nil, fmt.Errorf("algum slice está vazio")
	}
	slicefinal := make([]int, 0)
	for _, i := range a {
		for _, j := range b {
			if i == j {
				slicefinal = append(slicefinal, i)
			}
		}
	}
	return slicefinal, nil
}
