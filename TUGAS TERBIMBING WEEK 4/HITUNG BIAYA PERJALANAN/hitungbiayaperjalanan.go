package main

import "fmt"

func main() {
	var jarak, konsumsi, harga float64
	fmt.Print("Masukkan jarak dalam kilometer: ")
	fmt.Scan(&jarak)
	fmt.Print("Masukkan konsumsi bahan bakar kendaraan dalam kilometer per liter: ")
	fmt.Scan(&konsumsi)
	fmt.Print("Masukkan harga bahan bakar per liter: ")
	fmt.Scan(&harga)

	liter := jarak / konsumsi
	biaya := liter * harga

	fmt.Printf("Jadi, total biaya yang dikeluarkan adalah Rp%.0f. ", biaya)
}
