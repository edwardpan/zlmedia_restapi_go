package zlmedia

import (
	"context"
	"fmt"
)

// MediaAPI 流媒体管理相关API
type MediaAPI struct {
	client *Client
}

// NewMediaAPI 创建流媒体API实例
func NewMediaAPI(client *Client) *MediaAPI {
	return &MediaAPI{client: client}
}

// GetMediaAPI 获取流媒体API实例
func GetMediaAPI() *MediaAPI {
	return NewMediaAPI(GetClient())
}

// GetMediaListRequest 获取流列表请求参数
type GetMediaListRequest struct {
	Schema string `json:"schema,omitempty"` // 筛选协议，例如 rtsp或rtmp
	VHost  string `json:"vhost,omitempty"`  // 筛选虚拟主机，例如__defaultVhost__
	App    string `json:"app,omitempty"`    // 筛选应用名，例如 live
	Stream string `json:"stream,omitempty"` // 筛选流id，例如 test
}

// GetMediaList 获取流列表
// 获取ZLMediaKit中所有正在运行的流媒体列表
// 参数:
//   - Schema: 筛选协议，例如 rtsp或rtmp
//   - VHost: 筛选虚拟主机，例如__defaultVhost__
//   - App: 筛选应用名，例如 live
//   - Stream: 筛选流id，例如 test
//
// 返回: 流媒体列表信息
func (m *MediaAPI) GetMediaList(ctx context.Context, req *GetMediaListRequest) (*BaseResponse, error) {
	params := make(map[string]interface{})
	if req.Schema != "" {
		params["schema"] = req.Schema
	}
	if req.VHost != "" {
		params["vhost"] = req.VHost
	}
	if req.App != "" {
		params["app"] = req.App
	}
	if req.Stream != "" {
		params["stream"] = req.Stream
	}

	respBody, err := m.client.SendRequest(ctx, "GET", "/index/api/getMediaList", params)
	if err != nil {
		return nil, fmt.Errorf("获取流列表失败: %w", err)
	}

	return ParseResponse(respBody)
}

// CloseStreamRequest 关断单个流请求参数
type CloseStreamRequest struct {
	Schema string `json:"schema"`          // 协议，例如 rtsp或rtmp
	VHost  string `json:"vhost"`           // 虚拟主机，例如__defaultVhost__
	App    string `json:"app"`             // 应用名，例如 live
	Stream string `json:"stream"`          // 流id，例如 test
	Force  *bool  `json:"force,omitempty"` // 是否强制关闭(有人在观看是否还关闭)
}

// CloseStream 关断单个流
// 关闭指定的流媒体
// 参数:
//   - Schema: 协议，例如 rtsp或rtmp
//   - VHost: 虚拟主机，例如__defaultVhost__
//   - App: 应用名，例如 live
//   - Stream: 流id，例如 test
//   - Force: 是否强制关闭(有人在观看是否还关闭)
//
// 返回: 关闭结果
func (m *MediaAPI) CloseStream(ctx context.Context, req *CloseStreamRequest) (*BaseResponse, error) {
	params := map[string]interface{}{
		"schema": req.Schema,
		"vhost":  req.VHost,
		"app":    req.App,
		"stream": req.Stream,
	}
	if req.Force != nil {
		if *req.Force {
			params["force"] = "1"
		} else {
			params["force"] = "0"
		}
	}

	respBody, err := m.client.SendRequest(ctx, "GET", "/index/api/close_stream", params)
	if err != nil {
		return nil, fmt.Errorf("关闭流失败: %w", err)
	}

	return ParseResponse(respBody)
}

// CloseStreamsRequest 批量关断流请求参数
type CloseStreamsRequest struct {
	Schema string `json:"schema,omitempty"` // 协议，例如 rtsp或rtmp
	VHost  string `json:"vhost,omitempty"`  // 虚拟主机，例如__defaultVhost__
	App    string `json:"app,omitempty"`    // 应用名，例如 live
	Stream string `json:"stream,omitempty"` // 流id，例如 test
	Force  *bool  `json:"force,omitempty"`  // 是否强制关闭(有人在观看是否还关闭)
}

// CloseStreams 批量关断流
// 批量关闭符合条件的流媒体
// 参数:
//   - Schema: 协议，例如 rtsp或rtmp
//   - VHost: 虚拟主机，例如__defaultVhost__
//   - App: 应用名，例如 live
//   - Stream: 流id，例如 test
//   - Force: 是否强制关闭(有人在观看是否还关闭)
//
// 返回: 关闭结果
func (m *MediaAPI) CloseStreams(ctx context.Context, req *CloseStreamsRequest) (*BaseResponse, error) {
	params := make(map[string]interface{})
	if req.Schema != "" {
		params["schema"] = req.Schema
	}
	if req.VHost != "" {
		params["vhost"] = req.VHost
	}
	if req.App != "" {
		params["app"] = req.App
	}
	if req.Stream != "" {
		params["stream"] = req.Stream
	}
	if req.Force != nil {
		if *req.Force {
			params["force"] = "1"
		} else {
			params["force"] = "0"
		}
	}

	respBody, err := m.client.SendRequest(ctx, "GET", "/index/api/close_streams", params)
	if err != nil {
		return nil, fmt.Errorf("批量关闭流失败: %w", err)
	}

	return ParseResponse(respBody)
}
