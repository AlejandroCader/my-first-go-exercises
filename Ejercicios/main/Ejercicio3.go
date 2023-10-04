package main

import (
	"fmt"
)

func calcularSumaPares(inicio, fin int, c chan int) {
	suma := 0
	for i := inicio; i <= fin; i++ {
		if i%2 == 0 {
			suma += i
		}
	}
	c <- suma // Enviamos la suma de números pares al canal
}

func calcularSumaImpares(inicio, fin int, c chan int) {
	suma := 0
	for i := inicio; i <= fin; i++ {
		if i%2 != 0 {
			suma += i
		}
	}
	c <- suma // Enviamos la suma de números impares al canal
}

func main() {
	inicio := 1
	fin := 10
	canalPares := make(chan int)
	canalImpares := make(chan int)

	// Calculamos la suma de números pares e impares en goroutines separadas
	go calcularSumaPares(inicio, fin, canalPares)
	go calcularSumaImpares(inicio, fin, canalImpares)

	// Recibimos los resultados de las goroutines
	sumaPares := <-canalPares
	sumaImpares := <-canalImpares

	// Combinamos los resultados
	sumaTotal := sumaPares + sumaImpares

	fmt.Printf("Suma de números pares: %d\n", sumaPares)
	fmt.Printf("Suma de números impares: %d\n", sumaImpares)
	fmt.Printf("Suma total: %d\n", sumaTotal)
}
