package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("OPENAI_API_KEY is required")
	}

	testCases := []struct {
		name string
		url  string
		body string
	}{
		{
			name: "mj/submit/imagine",
			url:  "https://api.gpt.ge/mj/submit/imagine",
			body: `{"prompt":"A cute cat"}`,
		},
		{
			name: "v1/images/generations (midjourney)",
			url:  "https://api.gpt.ge/v1/images/generations",
			body: `{"model":"midjourney","prompt":"A cute cat"}`,
		},
		{
			name: "v1/chat/completions (midjourney)",
			url:  "https://api.gpt.ge/v1/chat/completions",
			body: `{"model":"midjourney","messages":[{"role":"user","content":"A cute cat"}]}`,
		},
	}

	for _, tc := range testCases {
		fmt.Println("Testing:", tc.name)
		req, _ := http.NewRequest("POST", tc.url, bytes.NewBuffer([]byte(tc.body)))
		req.Header.Set("Authorization", "Bearer "+apiKey)
		req.Header.Set("mj-api-secret", apiKey) // Sometimes MJ proxies use this
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
