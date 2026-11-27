package main
import "fmt"

func main(){
	var jam, biaya int

	fmt.Print("Masukkan jumlah jam: ")
	fmt.Scan(&jam)

	if jam <= 2 {
		biaya = 5000
	} else if jam >= 2 {
		biaya = 5000 + ((jam - 2) * 3000)
	} else {
		fmt.Print("Input tidak valid")
	}

	fmt.Print("Total biaya: ", biaya)
}