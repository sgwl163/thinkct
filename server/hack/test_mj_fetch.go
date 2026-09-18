package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	apiKey := "sk-S1rEFxwonvQR3wTx6d7d8a32C3E74c4389F1AeC7Fe7a1786"
	taskId := "1784915543848681" // use the task ID from the previous run
	
	url := "https://api.gpt.ge/mj/task/" + taskId + "/fetch"
	
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("mj-api-secret", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("Status: %d\nResponse: %s\n", resp.StatusCode, string(body))
}