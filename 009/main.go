package main

import "fmt"

func main() {
	a := 0
	b := 0
	c := 0
	s := 1000
	found := false
	for a = 1; a < s/3; a++ {
		for b = a; b < s/2; b++ {
			c = s - a - b
			if (a*a)+(b*b) == (c * c) {
				found = true
				break
			}
		}
		if found == true {
			break
		}
	}
	fmt.Printf("Los numeros son: %+v, %+v, %+v.\n", a, b, c)
	fmt.Printf("Producto de abc: %v.\n", a*b*c) // 31875000
}
