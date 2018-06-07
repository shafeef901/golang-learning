package main

import (
	"fmt"
	"router"
)

func main() {
	fmt.Println("Server ON!")

	e := router.New()

	e.Start(":8000")
}
