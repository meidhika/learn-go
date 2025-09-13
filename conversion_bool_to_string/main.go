package main

import (
	"fmt"
	"strconv"
)

func main() {
	truth := true
	str := strconv.FormatBool(truth)
	fmt.Println("Boolean to string ", str)

	val, _ := strconv.ParseBool("true")
	fmt.Println("String to boolean ", val)
}