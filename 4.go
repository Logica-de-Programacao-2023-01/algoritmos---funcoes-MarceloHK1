package marcelo

func segundo(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	maior := slice[0]
	second := slice[0]
	for _, i := range slice {
		if i > maior {
			maior = i
			second = maior
		} else if i > second && i != maior {
			second = i
		}
	}
	return second
}
