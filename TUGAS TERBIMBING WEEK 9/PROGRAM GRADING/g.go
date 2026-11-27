package main

import "fmt"

func main(){
	var jml_mahasiswa, a,b,c,d, nilai int

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jml_mahasiswa)
	a = 0
	b = 0
	c = 0
	d = 0

	for i := 1; i <= jml_mahasiswa; i++ {
		fmt.Printf("Masukkan nilai mahasiswa ke-%d: ", i)
		fmt.Scan(&nilai)
		if nilai >= 85 && nilai <= 100 {
			a += 1
		} else if nilai >= 75 {
			b += 1
		} else if nilai >= 60 {
			c += 1
		} else if nilai >= 40 {
			d += 1
		} else {
			fmt.Print("error")
		}
	}
	fmt.Println("TOTAL MAHASISWA UNTUK TIAP NILAI: ")
	fmt.Printf("A: %d\n", a)
	fmt.Printf("B: %d\n", b)
	fmt.Printf("C: %d\n", c)
	fmt.Printf("D: %d\n", d)
	
}