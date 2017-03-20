package main

import "fmt"

func main(){
	fmt.Print("Enter your name: ")
	var userName string
	fmt.Scan(&userName)
	fmt.Printf("Hello %v with name %v\n", userName, userName)
}