package main

import "fmt"

func main(){
	var jam, menit, detik int

	fmt.Print("Masukkan jam: ")
	fmt.Scan(&jam)
	fmt.Print("Masukkan menit: ")
	fmt.Scan(&menit)
	fmt.Print("Masukkan detik: ")
	fmt.Scan(&detik)

	jamkedetik := jam * 3600
	menitkedetik := menit * 60
	konversidetik := jamkedetik + menitkedetik + detik

	fmt.Printf("%d jam, %d menit, %d detik jika dikonversikan ke detik adalah %d detik.", jam, menit,detik,konversidetik)
}