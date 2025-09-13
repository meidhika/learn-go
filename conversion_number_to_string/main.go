package main

import (
	"fmt"
	"strconv"
)

func main(){
	var score int = 90
	var result string = strconv.Itoa(score) // int to string
	fmt.Println("Nilai awal: ",score)
	fmt.Println("Nilai akhir: ",result)
}