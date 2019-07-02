package main

import "fmt"

func main() {
	xText := "paraparaparadise"
	yText := "paragraph"

	X := nGram(xText, 2)
	Y := nGram(yText, 2)

	fmt.Printf("Union(X, Y) = %v\n", union(X, Y))
	fmt.Printf("Product(X, Y) = %v\n", product(X, Y))
	fmt.Printf("Difference(X, Y) = %v\n", difference(X, Y))
	fmt.Printf("'se' in X = %t, 'se' in Y = %t", in(X, "se"), in(Y, "se"))
}

func nGram(text string, n int) (result []string) {
	result = make([]string, len(text)+1-n)
	for i := 0; i < len(result); i++ {
		result[i] = string(text[i : i+n])
	}
	return
}

func union(x, y []string) []string {
	resultMap := make(map[string]int)
	for _, xEl := range x {
		resultMap[xEl] = 0
	}

	for _, yEl := range y {
		resultMap[yEl] = 0
	}
	return keys(resultMap)
}

func product(x, y []string) (result []string) {
	result = make([]string, 0)

	resultMap := make(map[string]int)
	for _, xEl := range x {
		if _, ok := resultMap[xEl]; !ok {
			resultMap[xEl] = 0
		}
	}

	for _, yEl := range y {
		if _, ok := resultMap[yEl]; ok {
			result = append(result, yEl)
			delete(resultMap, yEl)
		}
	}

	return
}

func difference(x, y []string) []string {
	resultMap := make(map[string]int)
	for _, xEl := range x {
		if _, ok := resultMap[xEl]; !ok {
			resultMap[xEl] = 0
		}
	}

	for _, yEl := range y {
		if _, ok := resultMap[yEl]; ok {
			delete(resultMap, yEl)
		}
	}

	return keys(resultMap)
}

func in(a []string, target string) bool {
	for _, el := range a {
		if el == target {
			return true
		}
	}
	return false
}

func keys(m map[string]int) []string {
	ks := []string{}
	for k, _ := range m {
		ks = append(ks, k)
	}
	return ks
}
