package main

import (
	"fmt"
)

var alphabet = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func modInverse(a, m int) int {
	for i := 1; i < m; i++ {
		if (a*i)%m == 1 {
			return i
		}
	}
	return 1
}

func encrypt(plainText string, a, b int) string {
	encryptedText := ""
	m := len(alphabet)
	for _, char := range plainText {
		if char == ' ' {
			encryptedText += " "
			continue
		}
		x := getIndex(char)
		encryptedChar := (a*x + b) % m
		encryptedText += string(alphabet[encryptedChar])
	}
	return encryptedText
}

func decrypt(cipherText string, a, b int) string {
	decryptedText := ""
	m := len(alphabet)
	aInverse := modInverse(a, m)
	for _, char := range cipherText {
		if char == ' ' {
			decryptedText += " "
			continue
		}
		x := getIndex(char)
		decryptedChar := (aInverse * (x - b + m)) % m
		decryptedText += string(alphabet[decryptedChar])
	}
	return decryptedText
}

func getIndex(char rune) int {
	for i, c := range alphabet {
		if c == char {
			return i
		}
	}
	return -1
}

func main() {
	plainText := "secretmessage"
	a, b := 5, 7
	encryptedText := encrypt(plainText, a, b)
	fmt.Println("Encrypted:", encryptedText)
	decryptedText := decrypt(encryptedText, a, b)
	fmt.Println("Decrypted:", decryptedText)
}
