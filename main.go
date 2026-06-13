package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("eDonish Pro v0.2.0")
	fmt.Println("Automated system for edonish.tj")
	
	if len(os.Args) < 3 {
		fmt.Println("Usage: edonish-pro <username> <password>")
		return
	}
	
	username := os.Args[1]
	password := os.Args[2]
	fmt.Printf("Login: %s\n", username)
	
	data := map[string]string{
		"username": username,
		"password": password,
	}
	jsonData, _ := json.Marshal(data)
	
	resp, err := http.Post("https://edonish.tj/auth/v1/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	
	result, _ := io.ReadAll(resp.Body)
	fmt.Println("Response:", string(result))
}
