package main

import "fmt"

func main() {
	// Cara klasik membuat variable

	var nama string = "Meidhika"
	var umur int = 20

	// cara menggunakan short declaration
	city := "Jakarta"
	year := 2022

	const pi = 3.14
	const name = "Meidhika"

	fmt.Println("Nama: ",nama)
	fmt.Println("Umur: ",umur)
	fmt.Println("City: ",city)
	fmt.Println("Year: ",year)
	fmt.Println("Pi	 : ",pi)
	fmt.Println("Name: ",name)

}