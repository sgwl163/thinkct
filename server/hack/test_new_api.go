package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	apiKey := "sk-S1rEFxwonvQR3wTx6d7d8a32C3E74c4389F1AeC7Fe7a1786"

	testCases := []struct {
		name string
		url  string
		body string
	}{
		{
			name: "images/generations (dall-e-3)",
			url:  "https://api.gpt.ge/v1/images/generations",
			body: `{"model":"dall-e-3","prompt":"A cute cat","n":1,"size":"1024x1024"}`,
		},
		{
			name: "chat/completions (gemini-3-pro-image)",
			url:  "https://api.gpt.ge/v1/chat/completions",
			body: `{"model":"gemini-3-pro-image","messages":[{"role":"user","content":"A cute cat"}]}`,
		},
	}

	for _, tc := range testCases {
		fmt.Println("Testing:", tc.name)
		req, _ := http.NewRequest("POST", tc.url, bytes.NewBuffer([]byte(tc.body)))
		req.Header.Set("Authorization", "Bearer "+apiKey)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("  Error:", err)
			continue
		}
		
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("  Status: %d\n  Response: %s\n\n", resp.StatusCode, string(body))
	}
}