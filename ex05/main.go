package main

import "fmt"

func main() {
	text := "I am an NLPer"
	biGram := nGram(text, 2)
	fmt.Printf("%v\n", biGram)
}

func nGram(text string, n int) (result []string) {
	result = make([]string, len(text)+1-n)
	for i := 0; i < len(result); i++ {
		result[i] = string(text[i : i+n])
	}
	return
}
