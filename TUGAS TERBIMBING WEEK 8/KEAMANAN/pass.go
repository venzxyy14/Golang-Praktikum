package main

import "fmt"

func main () {
	var usernamedata, usernameinput string
	var pwdata, pwinput int

	usernamedata = "Steven"
	pwdata = 123
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&usernameinput)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&pwinput)

	if usernameinput == usernamedata && pwdata == pwinput {
		fmt.Print("Pintu terbuka")
	} else {
		for i := 1; i <= 2; i++ {
			fmt.Print("Masukkan nama: ")
			fmt.Scan(&usernameinput)
			fmt.Print("Masukkan password: ")
			fmt.Scan(&pwinput)
			if usernameinput == usernamedata && pwdata == pwinput {
			fmt.Print("Pintu terbuka")
			break
			} else {
			fmt.Print("Tidak boleh masuk")
			fmt.Println("Silahkan coba lagi.")
			}
		}
		
}}