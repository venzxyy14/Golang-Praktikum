package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan angka untuk di faktorialkan: ")
	fmt.Scan(&n)

	faktorial := 1

	for i := 1; i <= n; i++ {
		faktorial *= i
	}

	fmt.Printf("%d! = %d",n, faktorial)
}