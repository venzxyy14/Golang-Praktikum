package main

import "fmt"

func main(){
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++{
		for j := 1; j <= n; j++{
			fmt.Printf("%d\t", i*j)
		}
	fmt.Println()
	}}

	// go run matrix.go