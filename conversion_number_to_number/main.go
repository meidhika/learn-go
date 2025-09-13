package main

import "fmt"

func main() {
	var num1 int = 10
	var num2 float64 = float64(num1)
	var num3 int = int(num2)
	fmt.Println("Nilai awal: ",num1)
	fmt.Println("Nilai akhir: ",num2)
	fmt.Println("Nilai akhir lagi: ",num3)
}