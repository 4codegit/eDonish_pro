package main

import (
	"fmt"
	"os"
	
	"encoding/json"
	"io"
	"net/http"
)

func main() {
	fmt.Println("eDonish Pro v0.2.0")
	
	if len(os.Args) < 2 {
		fmt.Println("Usage: edonish-pro <username>")
		return
	}
	
	username := os.Args[1]
	fmt.Println("User:", username)
	
	data := map[string]string{"username": username}
	body, _ := json.Marshal(data)
	
	resp, err := http.Post("https://edonish.tj/auth/v1/login", "application/json", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	
	result, _ := io.ReadAll(resp.Body)
	fmt.Println("Response:", string(result))
}
