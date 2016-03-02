package main

import "fmt"

func main() {
	numfibonaci := 0
	anterior := 0
	nuevo := 1
	suma := 0

	for numfibonaci < 10000000 {

		numfibonaci = anterior + nuevo
		if numfibonaci%2 != 0 && numfibonaci < 4000000 {
			fmt.Printf("%d,", numfibonaci)
			suma += numfibonaci
		}
		anterior = nuevo
		nuevo = numfibonaci
		//fmt.Printf("%d,", numfibonaci)
	}
	fmt.Printf("\nLa suma de los numeros: %d.\n", suma)
}
