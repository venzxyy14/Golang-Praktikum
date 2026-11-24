package main
import "fmt"

func main() {
	var x int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&x)

	if x > 0 {
		fmt.Print("Positif")
	} else {
		fmt.Print("Negatif")
	}
}