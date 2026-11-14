package main

import (
	"fmt"
)

const PI = 3.1415926535

func main() {
	var r int

	fmt.Print("Masukkan jari-jari: ")
	fmt.Scan(&r)

	rfloat := float32(r)
	keliling := 2 * PI * rfloat
	luas := PI * rfloat * rfloat

	fmt.Printf("Lingkaran dengan jari-jari %d memiliki luas %.2f dan keliling sebesar %.2f", r, keliling, luas)
}
