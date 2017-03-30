package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number to be halved: ")
	fmt.Scan(&num)
	fmt.Println(half(num))
}

func half(x int) (int, bool) {
	if x%2 == 0 {
		return x / 2, true
	} else {
		return x / 2, false
	}
}