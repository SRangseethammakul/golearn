package main

import "fmt"

func main() {
	langs := []string{"go", "python", "js"}
	fmt.Printf("langs : %#v\n", langs)

	a := langs[0:2] //[0,2)
	fmt.Printf("a : %#v\n", a)
	fmt.Printf("langs : %#v\n", langs)

	b := langs[1:3] //[1,3)
	fmt.Printf("b : %#v\n", b)
	fmt.Printf("langs : %#v\n", langs)

	n := langs[0:len(langs)]
	fmt.Printf("n : %#v\n", n)
	fmt.Println()
	r := langs[0:3]
	s := langs[:3]
	t := langs[0:]
	u := langs[:]
	fmt.Printf("r : %#v\n", r)
	fmt.Printf("s : %#v\n", s)
	fmt.Printf("t : %#v\n", t)
	fmt.Printf("u : %#v\n", u)
	printSlice(langs)
	cords := [4]string{"aa", "bb", "cc", "dd"}
	printSlice(cords[:]) //convert array to slices
}

func printSlice(ns []string) {
	fmt.Printf("print from slice : ns %#v\n", ns)
}
