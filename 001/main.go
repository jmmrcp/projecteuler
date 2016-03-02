/* If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. 
The sum of these multiples is 23. Find the sum of all the multiples of 3 or 5 below 1000.*/

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
	fmt.Printf("La suma de los numeros es: %d.\n", suma) // 233168
}
