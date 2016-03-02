package main

import (
	"fmt"
	"math"
)

// esPrimo no devuelve si es primo o no un numero.
//
// esPrimo uint -> bool
// esPrimo(11)
// true
func esPrimo(n uint) bool {
	var i uint
	for i = 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// esPrimo2 devuelve si es primo o no un numero. MAX RAPIDO
//
// esPrimo2 uint -> bool
// esPrimo(11)
// True
func esPrimo2(n uint) bool {
	var i uint
	f := float64(n)
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	if n < 9 {
		return true
	}
	if n%3 == 0 {
		return false
	}
	r := (math.Sqrt(f + 0.0)) + 1

	for i = 5; float64(i) <= r; i += 6 {
		if n%i == 0 {
			return false
		}
		if n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// genNumeros calcula el numero primo en la posicion n.
//
// genNumeros uint -> uint
// genNumeros(6)
// 13
func genNumeros(n uint) uint {
	var numero, contador, respuesta uint = 3, 1, 0
	for contador < n {
		if esPrimo(numero) {
			contador++
			respuesta = numero
		}
		numero++
	}
	return respuesta
}

func main() {
	fmt.Println(genNumeros(10001)) // 104743
}
