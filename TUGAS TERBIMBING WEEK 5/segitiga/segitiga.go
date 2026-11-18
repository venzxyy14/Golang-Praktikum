package main

import "fmt"

func main(){
	var n int

	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++{ // baris
		for j := 1; j <= i; j++{ // kolom
			fmt.Print("* ")
		}
	fmt.Println()
	}
}

// go run segitiga.go