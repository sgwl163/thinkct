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
	taskId := "1784915543848681" // 前面成功的 MJ 任务 ID

	testCases := []struct {
		name string
		url  string
		body string
	}{
		{
			name: "mj/submit/action (U1)",
			url:  "https://api.gpt.ge/mj/submit/action",
			body: fmt.Sprintf(`{"taskId":"%s","customId":"MJ::JOB::upsample::1::31307a65-856c-4b73-a621-0a6713833215"}`, taskId), // 只是测下格式
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
