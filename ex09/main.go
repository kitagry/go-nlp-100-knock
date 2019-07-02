package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	text := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	texts := strings.Split(text, " ")

	randomNum := make([]int, 0)
	textsCopy := make([]string, 0)
	for i := 1; i < len(texts)-1; i++ {
		if len(texts[i]) > 4 {
			randomNum = append(randomNum, i)
			textsCopy = append(textsCopy, texts[i])
		}
	}

	shuffle(textsCopy)

	for i := 0; i < len(randomNum); i++ {
		texts[randomNum[i]] = textsCopy[i]
	}

	fmt.Println(strings.Join(texts, " "))
}

func shuffle(nums []string) {
	for i := 0; i < len(nums); i++ {
		j := rand.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}
}
