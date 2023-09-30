package main

import (
	"fmt"
)

//func main() {
//	var alphabet = []rune{'а', 'б', 'в', 'г', 'д', 'е', 'ё', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я'}
//
//	word := "информатика"
//	key := "вася"
//	result := repeatAndTrim(key, word)
//
//	fmt.Println(result)
//
//	newStr := ""
//
//	for i, char := range word {
//		index := findIndex(alphabet, char)
//		keyIndex := findIndex(alphabet, rune(result[i]))
//		newIndex := (index + keyIndex) % len(alphabet)
//		newStr += string(alphabet[newIndex])
//	}
//
//	fmt.Println(newStr)
//}

func main() {
	//word := "информатика"
	key := "вася"
	fmt.Println(len(key))
	//fmt.Println(utf8.RuneCountInString(word))
	//result := repeatAndTrim(key, strconv.Itoa(utf8.RuneCountInString(word)))
	//fmt.Println(result)
}

func findIndex(arr []rune, char rune) int {
	for i, val := range arr {
		if val == char {
			return i
		}
	}
	return -1
}

func repeatAndTrim(s string, word string) string {
	repeated := ""
	length := len(word)
	for len(repeated) < length {
		repeated += s
	}
	return repeated[:length]
}
