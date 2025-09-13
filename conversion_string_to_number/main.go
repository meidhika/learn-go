package main

import (
	"fmt"
	"strconv"
)

func main() {
	var score string = "90"
	 result, err  := strconv.Atoi(score) // string to int

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}else{
		fmt.Println("Angka: ",result)
	}

}