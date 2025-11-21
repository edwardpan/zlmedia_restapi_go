package zlmedia

import (
	"context"
	"fmt"
)

// RecordAPI 录制管理相关API
type RecordAPI struct {
	client *Client
}

// NewRecordAPI 创建录制API实例
func NewRecordAPI(client *Client) *RecordAPI {
	return &RecordAPI{client: client}
}

// GetRecordAPI 获取录制API实例
func GetRecordAPI() *RecordAPI {
	return NewRecordAPI(GetClient())
}

// IsRecordingRequest 判断是否正在录制请求参数
type IsRecordingRequest struct {
	Type   int    `json:"type"`   // 0为hls，1为mp4
	VHost  string `json:"vhost"`  // 虚拟主机，例如__defaultVhost__
	App    string `json:"app"`    // 应用名，例如live
	Stream string `json:"stream"` // 流id，例如obs
}

// IsRecording 判断是否正在录制
// 检查指定流是否正在录制
// 参数:
//   - Type: 0为hls，1为mp4
//   - VHost: 虚拟主机，例如__defaultVhost__
//   - App: 应用名，例如live
//   - Stream: 流id，例如obs
//
// 返回: 录制状态信息
func (r *RecordAPI) IsRecording(ctx context.Context, req *IsRecordingRequest) (*BaseResponse, error) {
	params := map[string]interface{}{
		"type":   req.Type,
		"vhost":  req.VHost,
		"app":    req.App,
		"stream": req.Stream,
	}

	respBody, err := r.client.SendRequest(ctx, "GET", "/index/api/isRecording", params)
	if err != nil {
		return nil, fmt.Errorf("判断录制状态失败: %w", err)
	}

	return ParseResponse(respBody)
}

// StartRecordRequest 开始录制请求参数
type StartRecordRequest struct {
	Type           int    `json:"type"`                      // 0为hls，1为mp4
	VHost          string `json:"vhost"`                     // 虚拟主机，例如__defaultVhost__
	App            string `json:"app"`                       // 应用名，例如live
	Stream         string `json:"stream"`                    // 流id，例如obs
	CustomizedPath string `json:"customized_path,omitempty"` // 录制文件保存根目录，置空使用默认目录
	MaxSecond      *int   `json:"max_second,omitempty"`      // mp4录制切片大小，单位秒，置空时采用配置文件默认值
}

// StartRecord 开始录制
// 开始录制指定流
// 参数:
//   - Type: 0为hls，1为mp4
//   - VHost: 虚拟主机，例如__defaultVhost__
//   - App: 应用名，例如live
//   - Stream: 流id，例如obs
//   - CustomizedPath: 录制文件保存根目录，置空使用默认目录
//   - MaxSecond: mp4录制切片大小，单位秒，置空时采用配置文件默认值
//
// 返回: 开始录制结果
func (r *RecordAPI) StartRecord(ctx context.Context, req *StartRecordRequest) (*BaseResponse, error) {
	params := map[string]interface{}{
		"type":   req.Type,
		"vhost":  req.VHost,
		"app":    req.App,
		"stream": req.Stream,
	}

	if req.CustomizedPath != "" {
		params["customized_path"] = req.CustomizedPath
	}
	if req.MaxSecond != nil {
		params["max_second"] = *req.MaxSecond
	}

	respBody, err := r.client.SendRequest(ctx, "GET", "/index/api/startRecord", params)
	if err != nil {
		return nil, fmt.Errorf("开始录制失败: %w", err)
	}

	return ParseResponse(respBody)
}

// StopRecordRequest 停止录制请求参数
type StopRecordRequest struct {
	Type   int    `json:"type"`   // 0为hls，1为mp4
	VHost  string `json:"vhost"`  // 虚拟主机，例如__defaultVhost__
	App    string `json:"app"`    // 应用名，例如live
	Stream string `json:"stream"` // 流id，例如obs
}

// StopRecord 停止录制
// 停止录制指定流
// 参数:
//   - Type: 0为hls，1为mp4
//   - VHost: 虚拟主机，例如__defaultVhost__
//   - App: 应用名，例如live
//   - Stream: 流id，例如obs
//
// 返回: 停止录制结果
func (r *RecordAPI) StopRecord(ctx context.Context, req *StopRecordRequest) (*BaseResponse, error) {
	params := map[string]interface{}{
		"type":   req.Type,
		"vhost":  req.VHost,
		"app":    req.App,
		"stream": req.Stream,
	}

	respBody, err := r.client.SendRequest(ctx, "GET", "/index/api/stopRecord", params)
	if err != nil {
		return nil, fmt.Errorf("停止录制失败: %w", err)
	}

	return ParseResponse(respBody)
}

// GetMp4RecordFileRequest 获取录制文件夹内的文件列表请求参数
type GetMp4RecordFileRequest struct {
	VHost  string `json:"vhost"`  // 虚拟主机，例如__defaultVhost__
	App    string `json:"app"`    // 应用名，例如live
	Stream string `json:"stream"` // 流id，例如obs
	Period string `json:"period"` // 流的录制日期，格式为2020-02-01,如果不是完整的日期，那么是搜索录制文件夹列表，否则搜索对应日期下的mp4文件列表
}

// GetMp4RecordFile 获取录制文件夹内的文件列表
// 获取指定流的录制文件列表
// 参数:
//   - VHost: 虚拟主机，例如__defaultVhost__
//   - App: 应用名，例如live
//   - Stream: 流id，例如obs
//   - Period: 流的录制日期，格式为2020-02-01,如果不是完整的日期，那么是搜索录制文件夹列表，否则搜索对应日期下的mp4文件列表
//
// 返回: 录制文件列表
func (r *RecordAPI) GetMp4RecordFile(ctx context.Context, req *GetMp4RecordFileRequest) (*BaseResponse, error) {
	params := map[string]interface{}{
		"vhost":  req.VHost,
		"app":    req.App,
		"stream": req.Stream,
		"period": req.Period,
	}

	respBody, err := r.client.SendRequest(ctx, "GET", "/index/api/getMp4RecordFile", params)
	if err != nil {
		return nil, fmt.Errorf("获取录制文件列表失败: %w", err)
	}

	return ParseResponse(respBody)
}

// DeleteRecordDirectoryRequest 删除录制文件夹请求参数
type DeleteRecordDirectoryRequest struct {
	VHost  string `json:"vhost"`  // 虚拟主机，例如__defaultVhost__
	App    string `json:"app"`    // 应用名，例如live
	Stream string `json:"stream"` // 流id，例如obs
	Period string `json:"period"` // 流的录制日期，格式为2020-02-01,如果不是完整的日期，那么是删除录制文件夹，否则删除对应日期下的mp4文件
}

// DeleteRecordDirectory 删除录制文件夹
// 删除指定流的录制文件或文件夹
// 参数:
//   - VHost: 虚拟主机，例如__defaultVhost__
//   - App: 应用名，例如live
//   - Stream: 流id，例如obs
//   - Period: 流的录制日期，格式为2020-02-01,如果不是完整的日期，那么是删除录制文件夹，否则删除对应日期下的mp4文件
//
// 返回: 删除结果
func (r *RecordAPI) DeleteRecordDirectory(ctx context.Context, req *DeleteRecordDirectoryRequest) (*BaseResponse, error) {
	params := map[string]interface{}{
		"vhost":  req.VHost,
		"app":    req.App,
		"stream": req.Stream,
		"period": req.Period,
	}

	respBody, err := r.client.SendRequest(ctx, "GET", "/index/api/deleteRecordDirectory", params)
	if err != nil {
		return nil, fmt.Errorf("删除录制文件夹失败: %w", err)
	}

	return ParseResponse(respBody)
}

// GetSnapRequest 获取截图或生成实时截图请求参数
type GetSnapRequest struct {
	Url        string `json:"url"`         // 需要截图的 url，可以是本机的，也可以是远程主机的
	TimeoutSec int    `json:"timeout_sec"` // 截图失败超时时间，防止 FFmpeg 一直等待截图
	ExpireSec  int    `json:"expire_sec"`  // 截图的过期时间，该时间内产生的截图都会作为缓存返回
}

// GetSnap 获取截图或生成实时截图
// 获取截图或生成实时截图
// 参数:
//   - Url: 需要截图的 url，可以是本机的，也可以是远程主机的
//   - TimeoutSec: 截图失败超时时间，防止 FFmpeg 一直等待截图
//   - ExpireSec: 截图的过期时间，该时间内产生的截图都会作为缓存返回
//
// 返回: 截图结果
func (r *RecordAPI) GetSnap(ctx context.Context, req *GetSnapRequest) (*BaseResponse, error) {
	params := map[string]interface{}{
		"url":         req.Url,
		"timeout_sec": req.TimeoutSec,
		"expire_sec":  req.ExpireSec,
	}

	respBody, err := r.client.SendRequest(ctx, "GET", "/index/api/getSnap", params)
	if err != nil {
		return nil, fmt.Errorf("获取截图失败: %w", err)
	}

	return ParseResponse(respBody)
}
