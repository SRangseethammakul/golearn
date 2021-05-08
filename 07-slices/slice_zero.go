package main

import "fmt"

func main() {
	// var arr []string = []string{"python", "golang"}
	arr := []string{}
	fmt.Printf("arr : %#v\n", arr)
	if arr == nil {
		fmt.Println("Yes nil")
	} else {
		fmt.Printf("Not nil : %#v\n", arr)
	}
}
