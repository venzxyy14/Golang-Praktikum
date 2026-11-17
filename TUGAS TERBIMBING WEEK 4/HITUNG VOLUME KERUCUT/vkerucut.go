package main

import "fmt"

const PI = 3.1415926535

func main() {
	var n int
	var r, t, total_volume float64
	
	fmt.Print("Masukkan jumlah kerucut: ")
	fmt.Scan(&n)
	total_volume = 0

	for i := 0; i < n; i++ {
		fmt.Print("\nMasukkan jari-jari kerucut: ")
		fmt.Scan(&r)
		fmt.Print("Masukkan tinggi kerucut: ")
		fmt.Scan(&t)
		volume := (1.0 / 3.0) * PI * r * r * t
		fmt.Printf("Volume Kerucut: %.4f\n", volume)
		total_volume = total_volume + volume
	}

	fmt.Printf("Total Volume %d kerucut: %.4f", n, total_volume)
}