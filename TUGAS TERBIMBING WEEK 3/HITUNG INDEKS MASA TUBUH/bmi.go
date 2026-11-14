package main

import "fmt"

func main(){
	var berat, tinggi, bmi float64

	fmt.Print("Masukkan berat badan: ")
	fmt.Scan(&berat)
	fmt.Print("Masukkan tinggi badan: ")
	fmt.Scan(&tinggi)

	bmi = berat / (tinggi * tinggi)
	fmt.Printf("Jadi nilai BMI dari berat badan %.2fkg dan tinggi %.2fm adalah %.2f", berat, tinggi, bmi)
}