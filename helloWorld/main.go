package main

import "fmt"

func main() {
	fmt.Println("Hello Wolrd")
	if (true && false) || (false && true) || !(false && false){
		fmt.Println("Goodbye Wolrd")
	}
}
