package main
import "fmt"

func main(){
	var jumlah, nilai int

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jumlah)

	for i := 1; i <= jumlah; i++{
		fmt.Print("Masukkan nilai mahasiswa ke-", i, ": ")
		fmt.Scan(&nilai)
		if nilai >= 75 {
			fmt.Println("Lulus")
		} 
	}
}