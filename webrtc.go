package zlmedia

import (
	"context"
	"fmt"
)

// WebRTCAPI WebRTC管理相关API
type WebRTCAPI struct {
	client *Client
}

// NewWebRTCAPI 创建WebRTC API实例
func NewWebRTCAPI(client *Client) *WebRTCAPI {
	return &WebRTCAPI{client: client}
}

// GetWebRTCAPI 获取WebRTC API实例
func GetWebRTCAPI() *WebRTCAPI {
	return NewWebRTCAPI(GetClient())
}

// GetWebRTCApiRequest 获取WebRTC API请求参数
type GetWebRTCApiRequest struct {
	// 无额外参数，只需要secret
}

// GetWebRTCApi 获取WebRTC API
// 获取WebRTC相关的API信息
// 返回: WebRTC API信息
func (w *WebRTCAPI) GetWebRTCApi(ctx context.Context, req *GetWebRTCApiRequest) (*BaseResponse, error) {
	respBody, err := w.client.SendRequest(ctx, "GET", "/index/api/getWebRTCApi", nil)
	if err != nil {
		return nil, fmt.Errorf("获取WebRTC API失败: %w", err)
	}

	return ParseResponse(respBody)
}

// WebRTCRequest WebRTC请求参数
type WebRTCRequest struct {
	Api    string                 `json:"api"`              // WebRTC API类型，play或publish
	Type   string                 `json:"type"`             // 候选者类型，answer或offer
	SDP    string                 `json:"sdp"`              // SDP内容
	VHost  string                 `json:"vhost"`            // 虚拟主机，例如__defaultVhost__
	App    string                 `json:"app"`              // 应用名，例如live
	Stream string                 `json:"stream"`           // 流id，例如test
	Params map[string]interface{} `json:"params,omitempty"` // 其他参数
}

// WebRTC WebRTC接口
// 处理WebRTC的offer/answer交换
// 参数:
//   - Api: WebRTC API类型，play或publish
//   - Type: 候选者类型，answer或offer
//   - SDP: SDP内容
//   - VHost: 虚拟主机，例如__defaultVhost__
//   - App: 应用名，例如live
//   - Stream: 流id，例如test
//   - Params: 其他参数
//
// 返回: WebRTC响应信息
func (w *WebRTCAPI) WebRTC(ctx context.Context, req *WebRTCRequest) (*BaseResponse, error) {
	params := map[string]interface{}{
		"api":    req.Api,
		"type":   req.Type,
		"sdp":    req.SDP,
		"vhost":  req.VHost,
		"app":    req.App,
		"stream": req.Stream,
	}

	// 添加其他参数
	if req.Params != nil {
		for k, v := range req.Params {
			params[k] = v
		}
	}

	respBody, err := w.client.SendRequest(ctx, "POST", "/index/api/webrtc", params)
	if err != nil {
		return nil, fmt.Errorf("WebRTC请求失败: %w", err)
	}

	return ParseResponse(respBody)
}
