package main

import "fmt"

func main() {
	numero := 1000
	suma := 0
	for i := 0; i < numero; i++ {
		if i%3 == 0 || i%5 == 0 {
			suma += i
		} // End if
	} // End for
	fmt.Printf("La suma de los numeros es: %d.\n", suma)
}
