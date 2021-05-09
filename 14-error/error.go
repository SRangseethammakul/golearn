package main

import (
	"errors"
	"fmt"
	"log"
)

func ReadFile(name string) (string, error) {
	var err error = errors.New("file not found")
	return "", err
}

func main() {
	data, err := ReadFile("profile.txt")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("read file", data)
}
