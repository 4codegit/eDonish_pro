package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("eDonish Pro v0.2.0")
	if len(os.Args) >= 2 {
		fmt.Println("Args:", os.Args[1])
	}
}
