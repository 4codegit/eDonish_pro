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
	
	if len(os.Args) >= 2 {
		username := os.Args[1]
		fmt.Println("User:", username)
		
		reqBody, _ := json.Marshal(map[string]string{"username": username})
		resp, err := http.Post("https://edonish.tj/auth/v1/login", "application/json", nil)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()
		
		body, _ := io.ReadAll(resp.Body)
		var result map[string]interface{}
		json.Unmarshal(body, &result)
		
		fmt.Println("Response:", result)
	} else {
		fmt.Println("Usage: edonish-pro <username>")
	}
}
