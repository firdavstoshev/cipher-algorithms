package main

import "fmt"

var alphabet = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func main() {
	//fmt.Println(encrypt("secretmessage", 5, 7))
	fmt.Println(decrypt("tbrobypbtthlb", 5, 7))
}

func encrypt(word string, a, b int) string {
	encrypted := ""
	for _, char := range word {
		newIndex := (a*findIndex(alphabet, char) + b) % len(alphabet)
		newChar := alphabet[newIndex]
		encrypted += string(newChar)
	}
	return encrypted
}

func findIndex(arr []rune, char rune) int {
	for i, val := range arr {
		if val == char {
			return i
		}
	}
	return -1
}

func decrypt(word string, a, b int) string {
	decrypted := ""
	m := len(alphabet)
	for _, char := range word {
		newIndex := ((findIndex(alphabet, char) + m - b) / a) % m
		fmt.Print(newIndex, "\t")
		decrypted += string(alphabet[newIndex])
	}
	return decrypted
}
