package main

import (
	"fmt"
)

func main() {
	x := 12
	y := "気温"
	z := 22.4
	fmt.Println(template(x, y, z))
}

func template(x int, y string, z float64) string {
	return fmt.Sprintf("%d時の%sは%.1f", x, y, z)
}
