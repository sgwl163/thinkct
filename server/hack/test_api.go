package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://api.vectorengine.ai/v1/chat/completions"
	payload := []byte(`{"model":"gemini-3-pro-image","messages":[{"role":"user","content":"A futuristic city with flying cars"}]}`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Authorization", "Bearer sk-AKcoSZWXLytdtoCkWLyWqhfOThqVbZGjiMHy5oiOVIWSVis7")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}