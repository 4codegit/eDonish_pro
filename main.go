package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func login(username, password string) (*LoginResponse, error) {
	reqBody, _ := json.Marshal(LoginRequest{Username: username, Password: password})
	
	resp, err := http.Post("https://edonish.tj/auth/v1/login", "application/json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result LoginResponse
	json.Unmarshal(body, &result)
	
	return &result, nil
}

func main() {
	fmt.Println("eDonish Pro v0.2.0")
	fmt.Println("Console + GUI version")
	
	if len(os.Args) >= 3 {
		username := os.Args[1]
		password := os.Args[2]
		
		fmt.Printf("Logging in as %s...\n", username)
		result, err := login(username, password)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		
		if result.Success {
			fmt.Println("✅ Login successful!")
		} else {
			fmt.Printf("❌ Login failed: %s\n", result.Message)
		}
	} else {
		fmt.Println("Usage: edonish-pro <username> <password>")
		fmt.Println("GUI version: go run -tags gui main_gui.go")
	}
}
