package main

import "fmt"

// primefactor devuelve el mayor factor primo de un numero.
//
// primefactor uint -> uint
// primefactor(13192)
// 29
func primefactor(numero uint) uint {
	var i uint
	for numero != 1 {
		for i = 2; i < numero+1; i++ {
			if numero%i == 0 {
				break
			}
		}
		numero = numero / i
	}
	return i
}

func main() {
	var a uint = 600851475143
	fmt.Println(primefactor(a))
}
