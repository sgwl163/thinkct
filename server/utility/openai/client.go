package openai

import (
	"bufio"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// Client OpenAI 客户端
type Client struct {
	apiKey  string
	baseURL string
}

// NewClient 创建一个新的客户端
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseURL: "https://api.openai.com/v1", // 默认接口地址
	}
}

// SetBaseURL 设置自定义接口地址
func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

// Message 聊天消息体
type Message struct {
	Role    string      `json:"role"`
	Content interface{} `json:"content"`
}

// StreamOptions 流式请求附加选项
type StreamOptions struct {
	IncludeUsage bool `json:"include_usage"`
}

// ChatCompletionRequest 聊天请求体
type ChatCompletionRequest struct {
	Model         string         `json:"model"`
	Messages      []Message      `json:"messages"`
	Stream        bool           `json:"stream,omitempty"`
	StreamOptions *StreamOptions `json:"stream_options,omitempty"`
	Tools         interface{}    `json:"tools,omitempty"`
	Temperature   float32        `json:"temperature,omitempty"`
}

// Usage Token用量结构
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ChatCompletionResult 聊天请求完整结果
type ChatCompletionResult struct {
	Content string `json:"content"` // 回复内容
	Usage   *Usage `json:"usage"`   // Token用量统计
}

// ChatCompletion 发送聊天请求并获取回复
func (c *Client) ChatCompletion(ctx context.Context, req ChatCompletionRequest) (*ChatCompletionResult, error) {
	url := fmt.Sprintf("%s/chat/completions", c.baseURL)

	resp, err := g.Client().
		SetTimeout(10*time.Minute).
		SetHeader("Authorization", "Bearer "+c.apiKey).
		SetHeader("Content-Type", "application/json").
		Post(ctx, url, req)

	if err != nil {
		return nil, err
	}
	defer resp.Close()

	body := resp.ReadAllString()
	if resp.StatusCode != 200 {
		return nil, gerror.Newf("AI 接口请求失败: %s", body)
	}

	j, err := gjson.DecodeToJson(body)
	if err != nil {
		return nil, err
	}

	result := &ChatCompletionResult{
		Content: j.Get("choices.0.message.content").String(),
	}

	// 解析 Token 用量
	if j.Contains("usage") {
		var usage Usage
		if err := j.Get("usage").Scan(&usage); err == nil {
			result.Usage = &usage
		}
	}

	return result, nil
}

// ChatCompletionStream 发送流式聊天请求并获取回复
// 返回值: content (完整内容), usage (Token用量), error
func (c *Client) ChatCompletionStream(ctx context.Context, req ChatCompletionRequest, onChunk func(string) error) (string, *Usage, error) {
	req.Stream = true
	// 强制要求 OpenAI 在流式输出的最后一个 chunk 附带 usage 统计信息
	if req.StreamOptions == nil {
		req.StreamOptions = &StreamOptions{IncludeUsage: true}
	}
	url := fmt.Sprintf("%s/chat/completions", c.baseURL)

	resp, err := g.Client().
		SetTimeout(10*time.Minute).
		SetHeader("Authorization", "Bearer "+c.apiKey).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "text/event-stream").
		Post(ctx, url, req)

	if err != nil {
		return "", nil, err
	}
	defer resp.Close()

	if resp.StatusCode != 200 {
		return "", nil, gerror.Newf("AI 接口请求失败: %s", resp.ReadAllString())
	}

	var fullContent strings.Builder
	var lastUsage *Usage // 用于保存从最后一个 chunk 解析到的 Token 用量

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "data: ") {
			data := strings.TrimPrefix(line, "data: ")
			if data == "[DONE]" {
				break
			}
			j, err := gjson.DecodeToJson(data)
			if err != nil {
				continue // 忽略解析错误（如部分行或格式异常）
			}

			// 尝试解析 content 增量
			chunk := j.Get("choices.0.delta.content").String()
			if chunk != "" {
				fullContent.WriteString(chunk)
				if err := onChunk(chunk); err != nil {
					return fullContent.String(), lastUsage, err
				}
			}

			// 尝试解析 usage (OpenAI 配置了 stream_options.include_usage=true 后会在最后一个 chunk 返回 usage)
			if j.Contains("usage") && !j.Get("usage").IsNil() {
				var u Usage
				if err := j.Get("usage").Scan(&u); err == nil && u.TotalTokens > 0 {
					lastUsage = &u
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fullContent.String(), lastUsage, err
	}

	return fullContent.String(), lastUsage, nil
}
