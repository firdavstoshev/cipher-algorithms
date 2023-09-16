package main

import "fmt"

func main() {
	var firstAlphabet = []rune{'а', 'б', 'в', 'г', 'д', 'е', 'ё', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я'}
	var secondAlphabet = []rune{'й', 'ц', 'у', 'к', 'е', 'н', 'г', 'ш', 'щ', 'з', 'х', 'ъ', 'э', 'ж', 'д', 'л', 'о', 'р', 'п', 'а', 'в', 'ы', 'ф', 'я', 'ч', 'с', 'м', 'и', 'т', 'ь', 'б', 'ю', 'ё'}
	word := "криптография"
	var str string
	for _, let := range word {
		index := findIndex(firstAlphabet, let)
		if index == -1 {
			continue
		} else {
			str += string(secondAlphabet[index])
		}
	}
	fmt.Println(str)
}

func findIndex(arr []rune, char rune) int {
	for i, val := range arr {
		if val == char {
			return i
		}
	}
	return -1
}
