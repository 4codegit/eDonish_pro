package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("eDonish Pro v0.2.0")
	
	resp, err := http.Get("https://edonish.tj")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	
	body, _ := io.ReadAll(resp.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)
	
	fmt.Println("HTTP + JSON + IO works!")
}
