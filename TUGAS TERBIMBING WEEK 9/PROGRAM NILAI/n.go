package main

import "fmt"

func main(){
	var nilai float64

	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	if nilai >= 80 {
		fmt.Print("A")
	} else if nilai >= 70 {
		fmt.Print("B")
	} else if nilai >= 60 {
		fmt.Print("C")
	} else if nilai >= 50 {
		fmt.Print("D")
	} else {
		fmt.Print("E")
	}
}