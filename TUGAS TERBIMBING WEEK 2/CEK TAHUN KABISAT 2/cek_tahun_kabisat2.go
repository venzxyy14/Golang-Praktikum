package main

import "fmt"

func mainKabisat() {
	var tahun int
	var isKabisat bool

	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)
	isKabisat = (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)
	fmt.Printf("Kabisat: %t", isKabisat)

}
