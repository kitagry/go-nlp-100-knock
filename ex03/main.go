package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	text := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	re := regexp.MustCompile("[.|,]")
	text = re.ReplaceAllString(text, "")
	texts := strings.Split(text, " ")

	for i := 0; i < len(texts); i++ {
		fmt.Println(len(texts[i]))
	}
}
