package main

import "fmt"

func main() {
	var totalAwal, diskonPersen int

	fmt.Print("Masukkan total belanja awal: ")
	fmt.Scan(&totalAwal)
	fmt.Print("Masukkan besaran persentase diskon: ")
	fmt.Scan(&diskonPersen)
	nominalDiskon := (totalAwal * diskonPersen) / 100
	totalAkhir := totalAwal - nominalDiskon
	fmt.Printf("Jadi total belanja akhir yang harus dibayar adalah Rp%d", totalAkhir)
}
