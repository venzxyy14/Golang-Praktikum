package main
import "fmt"

func main(){
	var n int

	fmt.Print("Masukkan jumlah rumah: ")
	fmt.Scan(&n)

	nomor := -1

	for i := 1; i <= n; i ++{
		nomor = nomor + 2
		fmt.Println(nomor)
	}
}

