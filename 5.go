package marcelo

func achar(slice []int, a int) int {
	if len(slice) == 0 {
		return 0
	}
	for _, i := range slice {
		if i == a {
			break
		} else {
			return -1
		}
	}
	return a
}
