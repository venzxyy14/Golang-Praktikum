package main

import "fmt"

// func main(){
// 	var n int

// 	fmt.Print("Masukkan nilai: ")
// 	fmt.Scan(&n)

// 	for i := 1; i <= n; i++{
// 		d := i
// 		s := n % d == 0
// 		fmt.Print(d,s, "\n")            
// 	}
// }

func main(){
	var n, d int
	var s bool

	fmt.Scan(&n)
	
	for i := 1; i <= n; i++{
		d = i
		s = n % d == 0
		fmt.Println(d,s)
	}
}