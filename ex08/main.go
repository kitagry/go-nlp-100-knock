package main

import "fmt"

func main() {
	text := "aiue,asd12q38*()"
	cipheredText := cipher(text)
	fmt.Println(cipheredText)
	fmt.Println(cipher(cipheredText))
}

func cipher(text string) string {
	result := make([]rune, len(text))
	for i, t := range text {
		if ('A' <= t && t <= 'Z') || ('a' <= t && t <= 'z') {
			result[i] = 219 - t
		} else {
			result[i] = t
		}
	}
	return string(result)
}
