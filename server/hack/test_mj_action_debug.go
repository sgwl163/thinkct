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
	taskId := "1784915543848681" // 前面的任务ID

	testCases := []struct {
		name string
		url  string
		body string
	}{
		{
			name: "mj/submit/change (U1)",
			url:  "https://api.gpt.ge/mj/submit/change",
			body: fmt.Sprintf(`{"taskId":"%s","action":"U1"}`, taskId),
		},
		{
			name: "mj/submit/action (customId=U1)",
			url:  "https://api.gpt.ge/mj/submit/action",
			body: fmt.Sprintf(`{"taskId":"%s","customId":"U1"}`, taskId),
		},
	}

	for _, tc := range testCases {
		fmt.Println("Testing:", tc.name)
		req, _ := http.NewRequest("POST", tc.url, bytes.NewBuffer([]byte(tc.body)))
		req.Header.Set("Authorization", "Bearer "+apiKey)
		req.Header.Set("mj-api-secret", apiKey)
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
