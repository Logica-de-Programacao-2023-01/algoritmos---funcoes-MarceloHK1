package marcelo

func media(slice []int) float64 {
	if len(slice) == 0 {
		return 0
	}
	soma := 0
	for _, i := range slice {
		soma += i
	}

	return float64(soma) / float64(len(slice))
}

