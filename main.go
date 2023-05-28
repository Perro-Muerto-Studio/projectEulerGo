package main

import (
	"fmt"
	"projectEuler/problems"
)

func main() {
	fmt.Println("Resultado del problema 001:", problems.MultiplesOf3or5(1000))
	fmt.Println("Resultado del problema 002:", problems.EvenFibonacciNumbers(4_000_000))
	fmt.Println("Resultado del problema 003:", problems.LargestPrimeFactor(600_851_475_143))
}
