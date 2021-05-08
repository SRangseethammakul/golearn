package main

import "fmt"

func showAll(ns [4]string) {
	fmt.Printf("array show all : %#v\n", ns)
}
func main() {
	var arr = [4]string{"act1", "act2", "act3", "act4"}
	fmt.Printf("array : %#v\n", arr)
	n := arr[1]
	fmt.Printf("Array : %#v \n", n)

	arr[3] = "Suttipong"
	a := arr[3]
	fmt.Printf("Array : %#v \n", a)

	showAll(arr)
}
