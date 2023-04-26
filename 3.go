package marcelo

func junta(slice []string) string {
	if len(slice) == 0 {
		return ""
	}
	frase := ""
	for i := 0; i < len(slice); i++ {
		frase += slice[i]
		break
	}
	return frase
}
