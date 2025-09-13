package main

import (
	"fmt"
	"strings"
)

func main(){
	text := "Hello Dunia"

	// Mengubah menjadi huruf kecil
	fmt.Println("Lowercase:", strings.ToLower(text)) 

	// Mengubah menjadi huruf besar
	fmt.Println("Uppercase:", strings.ToUpper(text))

	// Mengechek apakah string dimulai dengan Hello
	fmt.Println("Starts with Hello:", strings.HasPrefix(text, "Hello"))

	// Mengechek apakah mengandung kata Dunia
	fmt.Println("Contains Dunia:", strings.Contains(text, "Dunia"))

	// Memisahkan string berdasarkan delimiter
	fruits := strings.Split("apel,cherry,banana", " ")
	fmt.Println("Split:", fruits)

	// Mengganti bagian string
	newText := strings.ReplaceAll(text, "Hello", "Hi", )
	fmt.Println("Replace:", newText)

	// Menghitung jumlah karakter
	fmt.Println("Length:", len(text))

	// Menghitung jumlah kata
	fmt.Println("Word Count:", len(strings.Split(text, " ")))

	// Menggabungkan string
	fmt.Println("Join:", strings.Join(fruits, "-"))

}