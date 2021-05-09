package main

import "fmt"

func main() {
	me := "act"
	fmt.Printf("i am %v\n", me)
	var addr *string = &me
	fmt.Println(addr)
	fmt.Printf("%T\n", addr)

	*addr = "Suttipong"
	fmt.Printf("i am %v\n", me)
}
