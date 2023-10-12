package main

import (
	"fmt"
	"projectEuler/problems"
)

func main() {
	fmt.Println("Resultado del problema 001: ", problems.MultiplesOf3or5(1000))
	fmt.Println("Resultado del problema 002: ", problems.EvenFibonacciNumbers(4_000_000))
	fmt.Println("Resultado del problema 003: ", problems.LargestPrimeFactor(600_851_475_143))
	fmt.Println("Resultado del problema 004: ", problems.LargestPalindromeProduct())
	fmt.Println("Resultado del problema 005: ", problems.SmallestMultiple(20))
	fmt.Println("Resultado del problema 006: ", problems.SumSquareDifference(100))
	fmt.Println("Resultado del problema 007: ", problems.XPrime(10_001))
	fmt.Println("Resultado del problema 008: ", problems.LargestProductInSeries(13))
}
