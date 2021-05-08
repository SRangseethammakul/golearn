package main

import "fmt"

func main() {
	num := 29
	if num%2 == 0 {
		fmt.Println("Even number")
	} else if num == 29 {
		fmt.Println("Lucky number")
	} else {
		fmt.Println("Odd number")
	}
}
