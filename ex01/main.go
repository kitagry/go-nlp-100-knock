package main

import "fmt"

func main() {
	text := "パタトクカシーー"
	textRunes := []rune(text)
	resultRunes := make([]rune, len(textRunes)/2)
	for i := 0; i < len(textRunes); i += 2 {
		resultRunes[i/2] = textRunes[i]
	}
	fmt.Println(string(resultRunes))
}
