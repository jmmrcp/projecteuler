package main

import "fmt"

// primos genera numeros primos
//
//primos -> int
// primos()
// 1
func primos(n int) {
	c := 1 // cantidad de numeros primos a generar
	p := 2 // primer numero primo
	d := 2 // numero para comparar
	for c <= n {
		if p%d == 0 {
			if p == d {
				fmt.Printf("%v, ", p)
				//return p
				c++
			}
			d = 2
			p++
		} else {
			d++ // iterador de numeros (numero /2 numero/3 numero/ ...)
		}
	}
	p--
	fmt.Printf("Valor de c: %v, p: %v, y d: %v", c, p, d)
}
func main() {
	primos(10)

}
