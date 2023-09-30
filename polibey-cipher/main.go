package main

import "fmt"

func main() {
	//var alphabet = []rune{'а', 'б', 'в', 'г', 'д', 'е', 'ё', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я', '.', '.', '.'}
	var arr [6][6]rune
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			arr[i][j] = 'a'
		}
	}
	fmt.Println(arr)
}
