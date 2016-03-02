package main

import "fmt"

func primefactor(numero uint) uint {
	var i uint
	for numero != 1 {
		for i = 2; i < numero+1; i++ {
			if numero%i == 0 {
				//fmt.Printf("%d / %d = %d.\n", numero, i, numero/i)
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
