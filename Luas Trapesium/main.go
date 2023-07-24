package main

import (
	"fmt"
	"math"
)

func alas_sisi_tinggi(s_atas, s_bawah, tinggi float64) float64 {
	return 0.5 * (s_atas + s_bawah) * tinggi
}

func sisi_alas_miring(s_atas, s_bawah, miring float64) float64 {
	// Menggunakan rumus luas trapesium tanpa tinggi
	// Tinggi dapat dihitung dengan menggunakan rumus Pythagoras
	tinggi_pyth := math.Sqrt(miring*miring - ((s_bawah - s_atas) * (s_bawah - s_atas)))
	return 0.5 * (s_atas + s_bawah) * tinggi_pyth
}

func main() {
	// Apakah tahu Tinggi Trapesium atau tidak
	var satu int
	fmt.Println("Apakah anda tahu tinggi trapesium?")
	fmt.Println("Masukkan (1) jika anda tahu tinggi trapesium")
	fmt.Println("Masukkan (0) jika anda tidak tahu tinggi trapesium")
	fmt.Print("Masukkan (1) atau (0) : ")
	fmt.Scan(&satu)

	if satu == 1 {

		// Kasus 1: Tinggi diketahui
		var s_atas, s_bawah, tinggi float64
		fmt.Print("Masukkan panjang sisi sejajar atas : ")
		fmt.Scan(&s_atas)
		fmt.Print("Masukkan panjang sisi sejajar bawah : ")
		fmt.Scan(&s_bawah)
		fmt.Print("Masukkan tinggi trapesium : ")
		fmt.Scan(&tinggi)

		luas := alas_sisi_tinggi(s_atas, s_bawah, tinggi)
		fmt.Printf("Luas trapesium dengan tinggi diketahui: %.2f\n", luas)
	} else if satu == 0 {

		// Kasus 2: Tinggi tidak diketahui
		var s_atas, s_bawah, miring float64
		fmt.Print("Masukkan panjang sisi sejajar atas : ")
		fmt.Scan(&s_atas)
		fmt.Print("Masukkan panjang sisi sejajar bawah : ")
		fmt.Scan(&s_bawah)
		fmt.Print("Masukkan panjang garis miring: ")
		fmt.Scan(&miring)

		luas := sisi_alas_miring(s_atas, s_bawah, miring)
		fmt.Printf("Luas trapesium dengan tinggi tidak diketahui: %.2f\n", luas)
	} else {
		fmt.Print("Terima kasih telah menggunakan program kami")
	}
}
