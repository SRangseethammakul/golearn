package main

import "fmt"

type User struct {
	UserName   string
	FullName   string
	ProfileURL string
}

func main() {
	u := User{
		UserName:   "act",
		FullName:   "Suttipong",
		ProfileURL: "www.google.com",
	} //initial struct
	// u.UserName = "act"
	// u.FullName = "Suttipong"
	// u.ProfileURL = "www.google.com"
	fmt.Printf("%#v\n", u)
	fmt.Println(u.FullName)
}
