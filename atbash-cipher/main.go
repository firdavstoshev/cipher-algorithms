package main

import "fmt"

var alphabet = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func main() {
	name := "firdavs"
	fmt.Println(encrypt(name)) // uriwzeh
}

func encrypt(word string) string {
	newWord := ""
	for _, char := range word {
		newIndex := len(alphabet) - 1 - findIndex(alphabet, char)
		newWord += string(alphabet[newIndex])
	}
	return newWord
}

func findIndex(arr []rune, char rune) int {
	for i, val := range arr {
		if val == char {
			return i
		}
	}
	return -1
}
