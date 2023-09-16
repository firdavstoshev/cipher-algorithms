package main

import (
	"fmt"
)

func main() {
	var arr = []rune{'.', ',', '-', ':', '"', '!', ';', 'А', 'Б', 'В', 'Г', 'Д', 'Е', ' ', 'Ж', 'З', 'И', 'Й', 'К', 'Л', 'М', 'Н', 'О', 'П', 'Р', 'С', 'Т', 'У', 'Ф', 'Х', 'Ц', 'Ч', 'Ш', 'Щ', 'Ъ', 'Ы', 'Ь', 'Э', 'Ю', 'Я'}
	cipheredText := "НКНКЫУШРМПЦУХУФРХЩЬЧЩШКМЭРМЬП.РМЫПЧПШРУРШКЫЩОЩМИ"

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
