package main

import "fmt"

type Phone interface {
	Call(number string)
}
type Apple struct {
	Name string
}

func (a Apple) Call(number string) { //call เป็น method ของ apple
	fmt.Println(a.Name, "calling", number)
}
func (a Apple) Answer() {
	fmt.Println(a.Name, "Answer")
}
func Dial(p Phone) {
	p.Call("44444444")
}

func main() {
	a := Apple{
		Name: "12ProMax",
	}
	Dial(a)
	a.Answer()
}
