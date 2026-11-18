package main

import "fmt"

func main(){
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	cek := 0

	for i := 0; i < n; i++{
		for j := 0; j < n; j++{
			cek = i + j
			if cek % 2 == 1 {
				fmt.Print(". ")
			} else {
				fmt.Print("# ")
			}
		}
	fmt.Println()
	}}

// go run catur.go