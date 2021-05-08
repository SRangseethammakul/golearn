package main

import "fmt"

// var name string = "Suttipong"
// var plus func(a, b int) int = func(a, b int) int { return a + b }

func say(str string) {
	fmt.Printf("my string : %s \n", str)
}
func cal(op func(int, int) int) {
	r := op(4, 5)
	fmt.Println(r)
}
func main() {
	var name = "Suttipong"
	fmt.Printf("Value : %v \n", name)
	fmt.Printf("Type : %T \n", name)

	var plus = func(a, b int) int { return a + b }
	p := plus(10, 5)
	fmt.Println(p)
	fmt.Printf("Type : %T \n", plus)
	cal(plus)

	var minus = func(a, b int) int { return a - b }
	cal(minus)

}
