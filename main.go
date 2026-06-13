package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("eDonish Pro v0.2.0")
	http.Get("https://edonish.tj")
	fmt.Println("Done")
}
