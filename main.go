package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("eDonish Pro v0.2.0")
	
	if len(os.Args) < 3 {
		fmt.Println("Usage: edonish-pro <username> <password>")
		return
	}
	
	username := os.Args[1]
	password := os.Args[2]
	fmt.Printf("Login: %s\n", username)
	
	data := map[string]string{"username": username, "password": password}
	body, _ := json.Marshal(data)
	
	resp, err := http.Post("https://edonish.tj/auth/v1/login", "application/json", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	
	result, _ := io.ReadAll(resp.Body)
	
	var respData map[string]interface{}
	json.Unmarshal(result, &respData)
	
	if respData["success"] == true {
		fmt.Println("Success!")
	} else {
		fmt.Println("Failed:", respData["message"])
	}
}
