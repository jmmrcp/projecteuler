package main

import "fmt"

func main() {
	i := 20 // numero minimo ... 20 % 20 = 1

	for i%11 != 0 ||
		i%12 != 0 ||
		i%13 != 0 ||
		i%14 != 0 ||
		i%15 != 0 ||
		i%16 != 0 ||
		i%17 != 0 ||
		i%18 != 0 ||
		i%19 != 0 ||
		i%20 != 0 {
		i += 20
	}
	fmt.Println("El numero es: ", i)

}
