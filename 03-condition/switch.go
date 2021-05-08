package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Monday
	switch today {
	case time.Monday:
		fmt.Println("today is Monday")
		fallthrough
	case time.Saturday, time.Sunday: // or
		fmt.Println("It's the Weekend")
	default:
		fmt.Printf("Hello %v \n", today)
	}
}
