package tqdm

func repeatChar(char rune, count int) string {
	slice := make([]rune, count)
	for i := range slice {
		slice[i] = char
	}
	return string(slice)
}
