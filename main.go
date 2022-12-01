package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("SOME_ENV"))
	fmt.Println("aaaa")
}
