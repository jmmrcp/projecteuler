package main

import (
	"fmt"
)

// sequence general la suma de numeros naturales
//
// sequence -> uint
// sequence()
// 1
// ...
// sequence()
// 28
func sequence() func() uint {
	var i, n uint
	return func() uint {
		i++
		n += i
		return n
	}
}

// divisores devuelve el n numeros por que es divisible un numero
//
// divisores uint -> uint
// divisores(10)
// 4
func divisores(n uint) uint {
	var i, r uint
	fmt.Printf("%2v: ", n)
	for i = 1; i <= n; i++ {
		if n%i == 0 {
			r++
		}
	}
	return r
}

// func calnumtriangular calcula el numero triangular de n
//
// calnumtriangular uint -> uint64
// calnumtriangular(12376)
// 76576500
func calnumtriangular(n uint) uint64 {
	var i uint
	var s uint64
	for i = 1; i < n; i++ {
		s += uint64(i)
	}
	return s
}
func calnumdivisores() {
	n := sequence()
	for {
		d := divisores(n())
		fmt.Println(d)
		if d > 500 {
			break
		}
	}
}

func main() {
	calnumdivisores()
}
