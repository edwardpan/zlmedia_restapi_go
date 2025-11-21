// Package zlmedia 实现基于ZLMediaKit REST API的操作模块
// ZLMediaKit是一个基于C++11的高性能运营级流媒体服务框架
// 支持WebRTC/RTSP/RTMP/HTTP/HLS/HTTP-FLV/WebSocket-FLV/HTTP-TS/HTTP-fMP4/WebSocket-TS/WebSocket-fMP4/GB28181/SRT等协议
package zlmedia

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Config ZLMediaKit客户端配置
type Config struct {
	BaseURL string        // ZLMediaKit API的基础URL，例如：http://localhost:80
	Secret  string        // API操作密钥(配置文件配置)
	Timeout time.Duration // HTTP客户端超时设置，默认为10秒
}

// Client ZLMediaKit客户端
type Client struct {
	config     Config
	httpClient *http.Client
}

// 全局ZLMediaKit客户端实例
var globalClient *Client

// InitClient 初始化ZLMediaKit客户端
// 参考文档: https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html
func InitClient(config Config) *Client {
	if globalClient != nil {
		return globalClient
	}

	// 设置默认超时
	if config.Timeout == 0 {
		config.Timeout = 10 * time.Second
	}

	// 确保baseURL不以/结尾
	config.BaseURL = strings.TrimSuffix(config.BaseURL, "/")

	globalClient = &Client{
		config: config,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	}

	return globalClient
}

// GetClient 获取全局ZLMediaKit客户端实例
func GetClient() *Client {
	if globalClient == nil {
		panic("ZLMediaKit客户端尚未初始化，请先调用InitClient")
	}
	return globalClient
}

// SendRequest 发送HTTP请求到ZLMediaKit API
func (c *Client) SendRequest(ctx context.Context, method, path string, params map[string]interface{}) ([]byte, error) {
	// 构建URL
	apiURL := fmt.Sprintf("%s%s", c.config.BaseURL, path)

	var reqBody io.Reader
	var contentType string

	if method == "GET" {
		// GET请求，参数放在URL中
		if params != nil {
			values := url.Values{}
			// 添加secret参数
			values.Set("secret", c.config.Secret)

			// 添加其他参数
			for key, value := range params {
				if value != nil {
					values.Set(key, fmt.Sprintf("%v", value))
				}
			}

			if len(values) > 0 {
				apiURL += "?" + values.Encode()
			}
		} else {
			// 只添加secret参数
			apiURL += "?secret=" + url.QueryEscape(c.config.Secret)
		}
	} else {
		// POST请求，参数放在body中
		if params != nil {
			// 添加secret参数到body
			params["secret"] = c.config.Secret

			jsonData, err := json.Marshal(params)
			if err != nil {
				return nil, fmt.Errorf("序列化请求体失败: %w", err)
			}
			reqBody = bytes.NewBuffer(jsonData)
			contentType = "application/json"
		} else {
			// 只发送secret参数
			values := url.Values{}
			values.Set("secret", c.config.Secret)
			reqBody = strings.NewReader(values.Encode())
			contentType = "application/x-www-form-urlencoded"
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, apiURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("创建HTTP请求失败: %w", err)
	}

	// 设置请求头
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	req.Header.Set("Accept", "application/json")

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送HTTP请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %w", err)
	}

	// 检查响应状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API请求失败，状态码: %d，响应: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// BaseResponse ZLMediaKit API的基础响应结构
type BaseResponse struct {
	Code int                    `json:"code"`           // 错误代码，0代表成功
	Msg  string                 `json:"msg,omitempty"`  // 不固定存在，可能为空
	Data map[string]interface{} `json:"data,omitempty"` // 返回数据，可能为空
}

// ParseResponse 解析ZLMediaKit API响应
func ParseResponse(respBody []byte) (*BaseResponse, error) {
	var response BaseResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	if response.Code != 0 {
		return &response, fmt.Errorf("API返回错误，代码: %d，消息: %s", response.Code, response.Msg)
	}

	return &response, nil
}
