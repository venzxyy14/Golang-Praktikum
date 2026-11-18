package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	kebalikan := 0

	for temp := n; temp > 0; temp /= 10 {
		digit := temp % 10
		kebalikan = (kebalikan * 10) + digit
	}

	isPalindrome := (n == kebalikan)
	fmt.Println(isPalindrome)
}