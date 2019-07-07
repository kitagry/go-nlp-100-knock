package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ikawaha/kagome/tokenizer"
)

func main() {
	t := tokenizer.New()

	data, err := ioutil.ReadFile("../neko.txt")
	if err != nil {
		log.Println(err)
		return
	}

	tokens := t.Tokenize(string(data))
	for _, token := range tokens {
		if token.Class == tokenizer.KNOWN && token.Features()[0] == "動詞" {
			fmt.Println(token.Surface)
		}
	}
}
