package main

import "fmt"

const PI = 3.1415926535

func main() {
	var r int

	fmt.Print("Masukkan nilai r: ")
	fmt.Scan(&r)
	rFloat := float64(r)

	volume := (4.0 / 3.0) * PI * rFloat * rFloat * rFloat
	luas := 4.0 * PI * rFloat * rFloat

	fmt.Printf("Bola dengan jejari %d memiliki volume %.3f dan luas kulit %.3f", r, volume, luas)
}
