package main

import "fmt"

func main() {
	var celcius int
	var fahreinheit int

	fmt.Print("Masukkan suhu dalam celcius: ")
	fmt.Scan(&celcius)

	fahreinheit = celcius*9/5 + 32
	fmt.Printf("Suhu: %d°F", fahreinheit)
}
