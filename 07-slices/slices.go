package main

import "fmt"

// func showAll(ns [4]string) {
// 	fmt.Printf("array show all : %#v\n", ns)
// }
func main() {
	arr := []string{"act1", "act2", "act3", "act4"}
	fmt.Printf("array : %#v\n", arr)
	n := arr[1]
	fmt.Printf("Array : %#v \n", n)

	arr[3] = "Suttipong"
	a := arr[3]
	fmt.Printf("Array : %#v \n", a)

	l := len(arr)
	fmt.Println(l)

	arr = append(arr, "act5", "act6", "act7")
	fmt.Printf("array : %#v\n", arr)
	fmt.Println(len(arr))
}
