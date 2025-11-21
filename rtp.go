package zlmedia

import (
	"context"
	"fmt"
)

// RTPAPI RTP管理相关API
type RTPAPI struct {
	client *Client
}

// NewRTPAPI 创建RTP API实例
func NewRTPAPI(client *Client) *RTPAPI {
	return &RTPAPI{client: client}
}

// GetRTPAPI 获取RTP API实例
func GetRTPAPI() *RTPAPI {
	return NewRTPAPI(GetClient())
}

// OpenRtpServerRequest 创建GB28181 RTP接收端口请求参数
type OpenRtpServerRequest struct {
	Port       int    `json:"port"`                  // 接收端口，0则为随机端口
	EnableTcp  *int   `json:"enable_tcp,omitempty"`  // 是否开启tcp模式，1为开启，0为关闭，默认为0
	StreamID   string `json:"stream_id"`             // 该端口绑定的流id
	ReUsePort  *int   `json:"re_use_port,omitempty"` // 是否重用端口，1为重用，0为不重用，默认为1
	SsrcFilter *int   `json:"ssrc_filter,omitempty"` // 是否开启ssrc过滤，1为开启，0为关闭，默认为0
}

// OpenRtpServer 创建GB28181 RTP接收端口
// 创建一个RTP接收端口，用于接收GB28181设备推送的RTP流
// 参数:
//   - Port: 接收端口，0则为随机端口
//   - EnableTcp: 是否开启tcp模式，1为开启，0为关闭，默认为0
//   - StreamID: 该端口绑定的流id
//   - ReUsePort: 是否重用端口，1为重用，0为不重用，默认为1
//   - SsrcFilter: 是否开启ssrc过滤，1为开启，0为关闭，默认为0
//
// 返回: 创建的RTP端口信息
func (rtp *RTPAPI) OpenRtpServer(ctx context.Context, req *OpenRtpServerRequest) (*BaseResponse, error) {
	params := map[string]interface{}{
		"port":      req.Port,
		"stream_id": req.StreamID,
	}

	if req.EnableTcp != nil {
		params["enable_tcp"] = *req.EnableTcp
	}
	if req.ReUsePort != nil {
		params["re_use_port"] = *req.ReUsePort
	}
	if req.SsrcFilter != nil {
		params["ssrc_filter"] = *req.SsrcFilter
	}

	respBody, err := rtp.client.SendRequest(ctx, "GET", "/index/api/openRtpServer", params)
	if err != nil {
		return nil, fmt.Errorf("创建RTP接收端口失败: %w", err)
	}

	return ParseResponse(respBody)
}

// CloseRtpServerRequest 关闭GB28181 RTP接收端口请求参数
type CloseRtpServerRequest struct {
	StreamID string `json:"stream_id"` // 调用openRtpServer接口时提供的流id
}

// CloseRtpServer 关闭GB28181 RTP接收端口
// 关闭指定的RTP接收端口
// 参数:
//   - StreamID: 调用openRtpServer接口时提供的流id
//
// 返回: 关闭结果
func (rtp *RTPAPI) CloseRtpServer(ctx context.Context, req *CloseRtpServerRequest) (*BaseResponse, error) {
	params := map[string]interface{}{
		"stream_id": req.StreamID,
	}

	respBody, err := rtp.client.SendRequest(ctx, "GET", "/index/api/closeRtpServer", params)
	if err != nil {
		return nil, fmt.Errorf("关闭RTP接收端口失败: %w", err)
	}

	return ParseResponse(respBody)
}

// ListRtpServerRequest 获取openRtpServer接口创建的所有RTP服务器请求参数
type ListRtpServerRequest struct {
	// 无额外参数，只需要secret
}

// ListRtpServer 获取openRtpServer接口创建的所有RTP服务器
// 获取所有RTP服务器的列表
// 返回: RTP服务器列表信息
func (rtp *RTPAPI) ListRtpServer(ctx context.Context, req *ListRtpServerRequest) (*BaseResponse, error) {
	respBody, err := rtp.client.SendRequest(ctx, "GET", "/index/api/listRtpServer", nil)
	if err != nil {
		return nil, fmt.Errorf("获取RTP服务器列表失败: %w", err)
	}

	return ParseResponse(respBody)
}

// StartSendRtpRequest 作为GB28181客户端，启动ps-rtp推流请求参数
type StartSendRtpRequest struct {
	VHost     string `json:"vhost"`                // 虚拟主机，例如__defaultVhost__
	App       string `json:"app"`                  // 应用名，例如live
	Stream    string `json:"stream"`               // 流id，例如test
	Ssrc      string `json:"ssrc"`                 // rtp推流的ssrc，ssrc不同时，可以推流到多个上级服务器
	DstURL    string `json:"dst_url"`              // 目标ip或域名
	DstPort   int    `json:"dst_port"`             // 目标端口
	IsUdp     *int   `json:"is_udp,omitempty"`     // 是否为udp模式，否则为tcp模式
	SrcPort   *int   `json:"src_port,omitempty"`   // 使用的本地端口，0则为随机端口
	Pt        *int   `json:"pt,omitempty"`         // 发送时，rtp的pt（uint8_t）,不传时默认为96
	UsePs     *int   `json:"use_ps,omitempty"`     // 发送时，rtp的负载类型。为1时，负载为ps；为0时，为es；不传时默认为1
	OnlyAudio *int   `json:"only_audio,omitempty"` // 当use_ps为0时，有效。为1时，发送音频；为0时，发送视频；不传时默认为0
}

// StartSendRtp 作为GB28181客户端，启动ps-rtp推流
// 启动RTP推流到指定的目标地址
// 参数:
//   - VHost: 虚拟主机，例如__defaultVhost__
//   - App: 应用名，例如live
//   - Stream: 流id，例如test
//   - Ssrc: rtp推流的ssrc，ssrc不同时，可以推流到多个上级服务器
//   - DstURL: 目标ip或域名
//   - DstPort: 目标端口
//   - IsUdp: 是否为udp模式，否则为tcp模式
//   - SrcPort: 使用的本地端口，0则为随机端口
//   - Pt: 发送时，rtp的pt（uint8_t）,不传时默认为96
//   - UsePs: 发送时，rtp的负载类型。为1时，负载为ps；为0时，为es；不传时默认为1
//   - OnlyAudio: 当use_ps为0时，有效。为1时，发送音频；为0时，发送视频；不传时默认为0
//
// 返回: 启动推流结果
func (rtp *RTPAPI) StartSendRtp(ctx context.Context, req *StartSendRtpRequest) (*BaseResponse, error) {
	params := map[string]interface{}{
		"vhost":    req.VHost,
		"app":      req.App,
		"stream":   req.Stream,
		"ssrc":     req.Ssrc,
		"dst_url":  req.DstURL,
		"dst_port": req.DstPort,
	}

	if req.IsUdp != nil {
		params["is_udp"] = *req.IsUdp
	}
	if req.SrcPort != nil {
		params["src_port"] = *req.SrcPort
	}
	if req.Pt != nil {
		params["pt"] = *req.Pt
	}
	if req.UsePs != nil {
		params["use_ps"] = *req.UsePs
	}
	if req.OnlyAudio != nil {
		params["only_audio"] = *req.OnlyAudio
	}

	respBody, err := rtp.client.SendRequest(ctx, "GET", "/index/api/startSendRtp", params)
	if err != nil {
		return nil, fmt.Errorf("启动RTP推流失败: %w", err)
	}

	return ParseResponse(respBody)
}

// StopSendRtpRequest 停止GB28181 ps-rtp推流请求参数
type StopSendRtpRequest struct {
	VHost  string `json:"vhost"`  // 虚拟主机，例如__defaultVhost__
	App    string `json:"app"`    // 应用名，例如live
	Stream string `json:"stream"` // 流id，例如test
	Ssrc   string `json:"ssrc"`   // rtp推流的ssrc
}

// StopSendRtp 停止GB28181 ps-rtp推流
// 停止指定的RTP推流
// 参数:
//   - VHost: 虚拟主机，例如__defaultVhost__
//   - App: 应用名，例如live
//   - Stream: 流id，例如test
//   - Ssrc: rtp推流的ssrc
//
// 返回: 停止推流结果
func (rtp *RTPAPI) StopSendRtp(ctx context.Context, req *StopSendRtpRequest) (*BaseResponse, error) {
	params := map[string]interface{}{
		"vhost":  req.VHost,
		"app":    req.App,
		"stream": req.Stream,
		"ssrc":   req.Ssrc,
	}

	respBody, err := rtp.client.SendRequest(ctx, "GET", "/index/api/stopSendRtp", params)
	if err != nil {
		return nil, fmt.Errorf("停止RTP推流失败: %w", err)
	}

	return ParseResponse(respBody)
}

// GetRtpInfoRequest 获取rtp推流信息请求参数
type GetRtpInfoRequest struct {
	StreamID string `json:"stream_id"` // 流id
}

// GetRtpInfo 获取rtp推流信息
// 获取指定流的RTP推流信息
// 参数:
//   - StreamID: 流id
//
// 返回: RTP推流信息
func (rtp *RTPAPI) GetRtpInfo(ctx context.Context, req *GetRtpInfoRequest) (*BaseResponse, error) {
	params := map[string]interface{}{
		"stream_id": req.StreamID,
	}

	respBody, err := rtp.client.SendRequest(ctx, "GET", "/index/api/getRtpInfo", params)
	if err != nil {
		return nil, fmt.Errorf("获取RTP推流信息失败: %w", err)
	}

	return ParseResponse(respBody)
}
