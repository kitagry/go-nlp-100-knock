package main

import "fmt"

func main() {
	text := "stressed"
	textSlice := []byte(text)
	for i := 0; i < len(textSlice)/2; i++ {
		tmp := textSlice[i]
		textSlice[i] = textSlice[len(textSlice)-1-i]
		textSlice[len(textSlice)-1-i] = tmp
	}
	fmt.Println(string(textSlice))
}
