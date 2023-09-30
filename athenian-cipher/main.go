package main

import "fmt"

func main() {
	var alphabet = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	a := 5
	b := 7
	word := "secretmessage"
	encrypted := ""
	for _, char := range word {
		newIndex := (a*findIndex(alphabet, char) + b) % len(alphabet)
		fmt.Print(newIndex, "\t")
		newChar := alphabet[newIndex]
		encrypted += string(newChar) + "\t"
	}
	fmt.Println("\n", encrypted)
}

func findIndex(arr []rune, char rune) int {
	for i, val := range arr {
		if val == char {
			return i
		}
	}
	return -1
}
