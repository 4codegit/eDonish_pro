package main

import (
	"encoding/json"
	"fmt"
)

type Test struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("eDonish Pro v0.2.0")
	t := Test{Name: "test"}
	json.Marshal(t)
	fmt.Println("JSON works!")
}
