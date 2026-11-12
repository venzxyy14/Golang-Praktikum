package main

import "fmt"

func main() {
	var c float64

	fmt.Print("Masukkan suhu dalam celcius: ")
	fmt.Scan(&c)

	f := (c * 9 / 5) + 32
	r := c * 4 / 5
	k := c + 273.15

	fmt.Println("==============================================================")
	fmt.Printf("BERIKUT DAFTAR PERUBAHAN SUHU DARI CELCIUS KE 4 SUHU LAINNYA: ")
	fmt.Println("")
	fmt.Println("==============================================================")
	fmt.Printf("Dalam Celcius: %.2f", c)
	fmt.Println("")
	fmt.Printf("Celcius Ke Fahrenheit: %.2f", f)
	fmt.Println("")
	fmt.Printf("Celcius Ke Reamur: %.2f", r)
	fmt.Println("")
	fmt.Printf("Celcius ke Kelvin: %.2f", k)

}
