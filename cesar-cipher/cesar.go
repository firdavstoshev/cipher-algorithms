package main

import (
	"fmt"
)

func main() {
	fmt.Println(cipher("фирдавс", 3))
	fmt.Println(decipher("члужгеф", 3))
}

func cipher(text string, shift int) string {
	var alphabet = []rune{'а', 'б', 'в', 'г', 'д', 'е', 'ё', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я'}
	var res string
	for _, char := range text {
		index := findIndex(alphabet, char)
		if index == -1 {
			continue
		} else {
			newIndex := (index + shift) % len(alphabet)
			res += string(alphabet[newIndex])
		}
	}
	return res
}

func decipher(text string, shift int) string {
	var alphabet = []rune{'а', 'б', 'в', 'г', 'д', 'е', 'ё', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я'}
	var res string
	for _, char := range text {
		index := findIndex(alphabet, char)
		if index == -1 {
			continue
		} else {
			newIndex := (index - shift + len(alphabet)) % len(alphabet)
			if newIndex < 0 {
				newIndex += len(alphabet)
			}
			res += string(alphabet[newIndex])
		}
	}
	return res
}

func findKeyByCipher(cipheredWord string) { // TODO
	for i := 1; i <= 33; i++ {
		var s string
		for _, char := range cipheredWord {
			s += string(char + rune(i))
		}
		fmt.Println(i, s)
	}
}

func findIndex(arr []rune, char rune) int {
	for i, val := range arr {
		if val == char {
			return i
		}
	}
	return -1
}
