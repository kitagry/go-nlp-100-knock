package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	text := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	re := regexp.MustCompile("[.|,]")
	text = re.ReplaceAllString(text, "")
	texts := strings.Split(text, " ")

	result := make(map[string]int)
	specialNums := []int{1, 5, 6, 7, 8, 9, 15, 16, 19}
	for i, t := range texts {
		if in(specialNums, i+1) {
			result[string(t[0])] = i
		} else {
			result[string(t[1])] = i
		}
	}
	fmt.Printf("%v", result)
}

func in(arr []int, target int) bool {
	for _, el := range arr {
		if el == target {
			return true
		}
	}
	return false
}
