package main

import "fmt"

func main() {
	const day = "monday"
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println("Day : ", day)
	fmt.Println("Sunday : ", Sunday)
	fmt.Println("Monday : ", Monday)
	fmt.Println("Tuesday : ", Tuesday)
	fmt.Println("Wednesday : ", Wednesday)
	fmt.Println("Thursday : ", Thursday)
	fmt.Println("Friday : ", Friday)
	fmt.Println("Saturday : ", Saturday)
}
