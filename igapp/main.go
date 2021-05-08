package main

import (
	"fmt"
	"github.com/srangseethammakul/igapp/time"
	"github.com/srangseethammakul/igapp/user"
)

func main() {
	user.Info("act", "suttipong", 24)
	t := time.Today()
	fmt.Printf("Value : %v \n", t)
}
