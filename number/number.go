package main

import "fmt"

func main(){
	// Signed Integer
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var i int = -100 // Ukuran tergantung arsitektur

	// Unsigned Integer
	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615
	var ui uint = 100 // Ukuran tergantung arsitektur


	// Deklarasi variable float32 dan float 64
	var f32 float32 = 33.1415926535
	var f64 float64 = 3.14159265358979323846

	// Menampilkan nilai
	fmt.Println("Floats :")
	fmt.Printf("float32 : %v\n", f32)
	fmt.Printf("float64 : %v\n", f64)

	// Menampilkan nilai
	fmt.Println("\nSigned Integers :")
	fmt.Printf("int8 : %v\n", i8)
	fmt.Printf("int16 : %v\n", i16)
	fmt.Printf("int32 : %v\n", i32)
	fmt.Printf("int64 : %v\n", i64)
	fmt.Printf("int : %v\n", i)


	fmt.Println("\nUnsigned Integers :")
	fmt.Printf("uint8 : %v\n", ui8)
	fmt.Printf("uint16 : %v\n", ui16)
	fmt.Printf("uint32 : %v\n", ui32)
	fmt.Printf("uint64 : %v\n", ui64)
	fmt.Printf("uint : %v\n", ui)




}