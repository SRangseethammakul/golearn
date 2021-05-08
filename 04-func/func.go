package main

import "fmt"

func info(name string, msg string) {
	fmt.Printf("My name is %v ,message : %v\n", name, msg)
}
func today() string {
	return "message"
}
func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	info("act", "hello")
	info("Suttipong", "sleep")
	day := today()
	fmt.Println(day)
	a, b := swap("act", "suttipong")
	fmt.Printf("A : %v, B : %v", a, b)
}
