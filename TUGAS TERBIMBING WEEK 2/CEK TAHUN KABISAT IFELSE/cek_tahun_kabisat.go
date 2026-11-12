package main

import "fmt"

func main() {
	var tahun int
	var isKabisat bool

	fmt.Print("Masukkan Tahun: ")
	fmt.Scan(&tahun)
	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		isKabisat = true
	} else {
		isKabisat = false
	}

	fmt.Printf("Kabisat: %t", isKabisat)

}
