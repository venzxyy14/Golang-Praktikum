package main
import "fmt"

func main() {
	var penghasilan int
	var nilai float32

	fmt.Print("Masukkan penghasilan orang tua: ")
	fmt.Scan(&penghasilan)
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	if nilai > 85 && penghasilan <= 3000000 {
		fmt.Print("Kamu dapat beasiswa")
	} else {
		fmt.Print("Tidak dapat beasiswa")
	}}