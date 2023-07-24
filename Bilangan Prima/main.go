package main

import (
	"fmt"
	"math"
)

func bilangan_prima(number int) bool {
	// Bilangan 1 dan angka negatif bukan bilangan prima
	if number <= 1 {
		return false
	}

	// Loop untuk memeriksa apakah ada pembagi selain 1 dan dirinya sendiri
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	// memasukkan nilai atau data
	var angka int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&angka)

	if bilangan_prima(angka) {
		fmt.Printf("%d adalah bilangan prima.\n", angka)
		fmt.Scanln()
	} else {
		fmt.Printf("%d bukan bilangan prima.\n", angka)
		fmt.Scanln()
	}
}
