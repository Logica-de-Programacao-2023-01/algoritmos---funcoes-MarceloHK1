package marcelo

func vogais(str string) int {
	if len(str) == 0 {
		return 0
	}
	vowel := "aeiouAEIOU"
	quantia := 0
	for _, letra := range str {
		for _, vogal := range vowel {
			if letra == vogal {
				quantia++
				break
			}
		}
	}
	return quantia
}
