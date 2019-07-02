package main

import "fmt"

func main() {
	policeCar := []rune("パトカー")
	taxi := []rune("タクシー")
	resultRunes := make([]rune, len(policeCar)+len(taxi))
	for i := 0; i < len(resultRunes); i++ {
		if i%2 == 0 {
			resultRunes[i] = policeCar[i/2]
		} else {
			resultRunes[i] = taxi[i/2]
		}
	}
	fmt.Println(string(resultRunes))
}
