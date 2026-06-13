package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("eDonish Pro")
	if len(os.Args) > 1 {
		fmt.Println("Arg:", os.Args[1])
		http.Get("https://edonish.tj")
	}
}
