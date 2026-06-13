package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("eDonish Pro")
	if len(os.Args) > 1 {
		fmt.Println("Arg:", os.Args[1])
	}
}
