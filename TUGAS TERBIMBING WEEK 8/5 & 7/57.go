package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&n)

	if (n%5 == 0) || (n%7 == 0) {
		fmt.Print("Memenuhi syarat")
	} else {
		fmt.Print("Tidak memenuhi syarat")
	}
}
