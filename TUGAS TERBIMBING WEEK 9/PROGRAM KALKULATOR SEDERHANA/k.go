package main
import "fmt"

func main(){
	var op string
	var a,b int

	fmt.Print("Masukkan pilihan operasi hitung (ex: tambah, kurang, kali, bagi): ")
	fmt.Scan(&op)
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scan(&b)

	if op == "tambah" {
		fmt.Print(a + b)
	} else if op == "kurang" {
		fmt.Print(a - b)
	} else if op == "kali" {
		fmt.Print(a * b)
	} else if op == "bagi" {
		fmt.Print(a / b)
	} else {
		fmt.Print("error")
	}
}