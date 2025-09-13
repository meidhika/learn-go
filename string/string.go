package main

import "fmt"

func main(){
	// Deklarasi string dengan kutip ganda
	nama := "Meidhika"
	pesan := "Selamat Datang di rumah saya"

	// Deklarasi string dengan kutip satu
	
	paragraf := `Meidhika Nawa Sapta disini,
	terima kasih telah berkunjung di rumah saya`

	fmt.Println("Nama: ",nama)
	fmt.Println("Pesan: ",pesan)
	fmt.Println("Paragraf: \n",paragraf)
}