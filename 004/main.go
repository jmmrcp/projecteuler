package main

import "fmt"

func invertirnumero(numero int) int {
	invertido := 0
	for numero > 0 {
		invertido = 10*invertido + numero%10
		numero = numero / 10
	}
	return invertido
}

func espalindromo(numero int) bool {
	return numero == invertirnumero(numero)
}

func generador() int {
	maximo := 0
	i := 999
	j := 0
	for i >= 100 {
		j = 999
		for j >= i {
			if i*j <= maximo {
				break
			}
			if espalindromo(i*j) == true {
				maximo = i * j
			}
			j--
		}
		i--
	}
	return maximo
}

func main() {
	numero := generador()
	fmt.Println(numero)
}
