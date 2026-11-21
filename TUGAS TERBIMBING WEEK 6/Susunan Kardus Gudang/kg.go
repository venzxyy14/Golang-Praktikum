package main

import "fmt"

func main(){
	var n int

	fmt.Print("Masukkan jumlah kardus di baris paling atas: ")
	fmt.Scan(&n)

	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++{
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}