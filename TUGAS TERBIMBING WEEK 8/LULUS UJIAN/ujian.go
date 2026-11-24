package main
import "fmt"

func main() {
	var nilai int

	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	if nilai >= 60 && nilai <= 100 {
		fmt.Print("Selamat anda dinyatakan lulus!")
	} else {
		fmt.Print("Anda tidak lulus, tetap semangat!")
	}
}