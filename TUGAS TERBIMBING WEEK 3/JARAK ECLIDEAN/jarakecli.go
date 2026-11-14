package main

import (
	"fmt"
	"math"
)

func main(){
	var x1, y1, x2, y2 float64

	fmt.Print("Masukkan nilai x1: ")
	fmt.Scan(&x1)
	fmt.Print("Masukkan nilai y1: ")
	fmt.Scan(&y1)
	fmt.Print("Masukkan nilai x2: ")
	fmt.Scan(&x2)
	fmt.Print("Masukkan nilai y2: ")
	fmt.Scan(&y2)

	dx := x2 - x1
	dy := y2 - y1 

	jaraklurus := math.Sqrt(dx*dx + dy*dy)

	fmt.Printf("Jadi jarak lurus antara x1 = %.1f, y1 = %.1f, x2 = %.1f dan y2 = %.1f adalah %.1f", x1, y1, x2, y2, jaraklurus)
}