package main

import "fmt"

func main() {
	fmt.Println(cipher("криптография", 3))
}

func cipher(text string, key int) string {
	var res string
	for _, letter := range text {
		res += string(letter + int32(key))
	}
	return res
}

func decipher(text string, key int) string {
	var res string
	for _, letter := range text {
		res += string(letter - int32(key))
	}
	return res
}

func findKeyByCipher(cipheredWord string) {
	for i := 1; i <= 33; i++ {
		var s string
		for _, let := range cipheredWord {
			s += string(let + int32(i))
		}
		fmt.Println(i, s)
	}
}
