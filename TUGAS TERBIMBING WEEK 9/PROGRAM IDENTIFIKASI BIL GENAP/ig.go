package main
import "fmt"

func main(){
	var jumlah, nomor int

	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&jumlah)

	for i := 1; i <= jumlah; i++ {
		fmt.Printf("Masukkan nomor peserta ke-%d: ", i)
		fmt.Scan(&nomor)
		if nomor % 2 == 0 {
			fmt.Println(nomor)
		}
	}
}