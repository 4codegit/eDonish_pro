package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func main() {
	fmt.Println("eDonish Pro v0.2.0")
	fmt.Println("Automated system for edonish.tj")
	
	if len(os.Args) < 3 {
		fmt.Println("Usage: edonish-pro <username> <password>")
		fmt.Println("Example: ./edonish-pro 200117707 test123")
		return
	}
	
	username := os.Args[1]
	password := os.Args[2]
	
	fmt.Printf("Logging in as %s...\n", username)
	
	data := map[string]string{"username": username, "password": password}
	body, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON Error:", err)
		return
	}
	
	resp, err := http.Post("https://edonish.tj/auth/v1/login", "application/json", nil)
	if err != nil {
		fmt.Println("HTTP Error:", err)
		return
	}
	defer resp.Body.Close()
	
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Error:", err)
		return
	}
	
	var loginResp LoginResponse
	json.Unmarshal(result, &loginResp)
	
	if loginResp.Success {
		fmt.Println("✅ Login successful!")
	} else {
		fmt.Printf("❌ Login failed: %s\n", loginResp.Message)
	}
}
