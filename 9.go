package marcelo

import (
	"fmt"
	"strings"
)

func nove(a string) (b []string, err error) {
	if len(a) == 0 {
		return nil, fmt.Errorf("a string nao existe")
	}
	slicefinal := strings.Split(a, " ")
	return slicefinal, nil
}
