package main

import "fmt"

func main(){
	var x,y float64

	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan y: ")
	fmt.Scan(&y)

	if x > 0 && y > 0 {
		fmt.Print("KUADRAN I")
	} else if x < 0 && y > 0  {
		fmt.Print("KUADRAN II")
	} else if x < 0 && y < 0 {
		fmt.Print("KUADRAN III")
	} else if x > 0 && y < 0 {
		fmt.Print("KUADRAN IV")
	} else {
		fmt.Print("ERROR")
	}
}