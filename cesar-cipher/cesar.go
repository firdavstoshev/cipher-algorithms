package main

import (
	"fmt"
)

func main() {
	var arr = []rune{'А', 'Б', 'В', 'Г', 'G', 'Д', 'Е', '~', 'Ж', 'З', 'И', 'm', 'Й', 'К', 'Q', 'Л', 'М', 'Н', 'О', 'П', 'Р', 'С', 'Т', 'У', 'U', 'Ф', 'Х', 'H', 'Ч', 'J', 'Ш', 'Ъ', 'Э', 'Ю', 'Я'}
	cipheredText := "ОЛИП HХ ЗШЪ"

	for shift := 1; shift < len(arr); shift++ {
		var decryptedText string
		for _, char := range cipheredText {
			if char == '.' {
				decryptedText += " "
			} else {
				index := findIndex(arr, char)
				if index != -1 {
					shiftedIndex := (index - shift + len(arr)) % len(arr)
					decryptedText += string(arr[shiftedIndex])
				} else {
					decryptedText += string(char)
				}
			}
		}
		fmt.Printf("Key %d: %s\n", shift, decryptedText)
	}
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
