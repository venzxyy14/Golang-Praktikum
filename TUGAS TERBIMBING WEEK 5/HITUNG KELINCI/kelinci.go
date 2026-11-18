package main

import "fmt"

func main(){
	var n int

	fmt.Print("Masukkan jumlah bulan yang ingin dihitung: ")
	fmt.Scan(&n)

	a := 0
	b := 1

	for i := 1; i <= n; i++{
		fmt.Print(a, " ")
		next := a + b
		a = b
		b = next
	}
	fmt.Println()
}