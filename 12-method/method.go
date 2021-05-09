package main

import "fmt"

type User struct { //lower case private and upper case public
	UserName   string
	FullName   string
	ProfileURL string
}

// func info(u User) {
// 	fmt.Println(u.FullName)
// 	fmt.Println(u.UserName)
// 	fmt.Println(u.ProfileURL)
// }
func (u User) Info() { //method
	fmt.Println(u.FullName)
	fmt.Println(u.UserName)
	fmt.Println(u.ProfileURL)
}

func main() {
	u := User{
		UserName:   "act",
		FullName:   "Suttipong",
		ProfileURL: "www.google.com",
	}
	// info(u)
	u.Info()
}
