package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Masukkan angka pertama: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scan(&b)

	hasil := 0 

	for i := 1; i <= b; i++ {
		hasil += a
	}
	fmt.Printf("Hasil perkalian antara %d dan %d adalah %d", a, b, hasil)
}