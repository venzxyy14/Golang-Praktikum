package main

import "fmt"

func main(){
	var jml_musuh, jml_koin, bonus int
	fmt.Print("Masukkan jumlah musuh yang dikalahkan: ")
	fmt.Scan(&jml_musuh)
	fmt.Print("Masukkan jumlah koin yang dikumpulkan: ")
	fmt.Scan(&jml_koin)
	fmt.Print("Masukkan bonus waktu: ")
	fmt.Scan(&bonus)

	total_skor := (jml_musuh * 100) + (jml_koin * 25) + bonus
	fmt.Printf("Total Skor: %d", total_skor)
}