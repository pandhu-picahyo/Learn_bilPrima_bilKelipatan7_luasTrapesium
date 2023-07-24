package main

import "fmt"

func Bil_Kelipatan_7(number int) bool {
	// Cek apakah angka tersebut merupakan kelipatan 7
	if number%7 == 0 {
		return true
	}
	return false
}

func main() {
	var angka int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&angka)

	if Bil_Kelipatan_7(angka) {
		fmt.Printf("%d adalah kelipatan 7.\n", angka)
	} else {
		fmt.Printf("%d bukan kelipatan 7.\n", angka)
	}
}
