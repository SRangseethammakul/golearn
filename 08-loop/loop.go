package main

import "fmt"

func main() {
	//for loop
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	//while loop
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//loop for  ever
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }

	langs := []string{"go", "python", "js"}
	for i := 0; i < len(langs); i++ {
		value := langs[i]
		fmt.Printf("index : %d, value : %s\n", i, value)
	}

	fmt.Println("for range slice")
	for index, value := range langs {
		fmt.Println(index, ":", value)
	}
	fmt.Println("for range slice value only")
	for _, value := range langs {
		fmt.Println(value)
	}
	fmt.Println("for range slice index only")
	for index := range langs {
		fmt.Println(index)
	}
}
