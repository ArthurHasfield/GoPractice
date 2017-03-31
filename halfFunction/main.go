package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number to be halved: ")
	fmt.Scan(&num)
	fmt.Println(half(num))
}

func half(x int) (float64, bool) {
	if x%2 == 0 {
		return float64(x) / 2, true
	} else {
		return float64(x) / 2, false
	}
}