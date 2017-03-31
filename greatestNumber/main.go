package main

import "fmt"

func greatestNum(numList ...int) int {
	var currentNum int
	for i := 0; i < len(numList); i++{
		if numList[i] > currentNum {
			currentNum = numList[i]
		}
	}
	return currentNum
}

func main (){
	var num = greatestNum(8, 3568, 17, 6994, 2015, 1006, 750)
	fmt.Println(num)
}