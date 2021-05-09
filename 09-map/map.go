package main

import "fmt"

func main() {
	//key value
	status := map[int]string{
		200: "OK",
		400: "Bad Request",
		500: "Server Error",
	}
	fmt.Printf("% #v\n", status)
	l := len(status)
	fmt.Println(l)

	status[200] = "Success"
	status[204] = "No Content"
	delete(status, 204)
	fmt.Printf("% #v\n", status)

	v, ok := status[444]
	if ok {
		fmt.Printf("value : %q \n\n", v)
	} else {
		fmt.Println("not found")
	}

	if v, ok := status[55]; ok {
		fmt.Printf("value : %q \n\n", v)
	} else {
		fmt.Println("not found")
	}

	var m = make(map[string]string)
	if m == nil {
		fmt.Println("Yes nil")
	} else {
		fmt.Printf("Not nil : %#v\n", m)
	}

}
