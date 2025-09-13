package make

import "fmt"

func main(){
	// Contoh Tipe data boolean
	var isTrue bool = true
	var isFalse bool = false

	isLoggedIn := true
	hasPermission := false

	// Menampilkan nilai Boolean
	fmt.Println("\nBoolean :")
	fmt.Printf("isTrue ? %v\n", isTrue)
	fmt.Printf("isFalse ? %v\n", isFalse)
	fmt.Printf("isLoggedIn ? %v\n", isLoggedIn)
	fmt.Printf("hasPermission ? %v\n", hasPermission)
}