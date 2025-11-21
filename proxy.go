package zlmedia

import (
	"context"
	"fmt"
)

// ProxyAPI 代理管理相关API
type ProxyAPI struct {
	client *Client
}

// NewProxyAPI 创建代理API实例
func NewProxyAPI(client *Client) *ProxyAPI {
	return &ProxyAPI{client: client}
}

// GetProxyAPI 获取代理API实例
func GetProxyAPI() *ProxyAPI {
	return NewProxyAPI(GetClient())
}

// AddStreamProxyRequest 添加拉流代理请求参数
type AddStreamProxyRequest struct {
	VHost         string   `json:"vhost"`                     // 添加的流的虚拟主机，例如__defaultVhost__
	App           string   `json:"app"`                       // 添加的流的应用名，例如live
	Stream        string   `json:"stream"`                    // 添加的流的id名，例如test
	URL           string   `json:"url"`                       // 拉流地址，例如rtmp://live.hkstv.hk.lxdns.com/live/hks2
	RtpType       *int     `json:"rtp_type,omitempty"`        // rtsp拉流时，拉流方式，0：tcp，1：udp，2：组播
	TimeoutSec    *float64 `json:"timeout_sec,omitempty"`     // 拉流超时时间，单位秒，float类型
	RetryCount    *int     `json:"retry_count,omitempty"`     // 拉流重试次数,不传此参数或传值<=0时，则无限重试
	EnableHLS     *bool    `json:"enable_hls,omitempty"`      // 是否转hls-ts
	EnableHLSFmp4 *bool    `json:"enable_hls_fmp4,omitempty"` // 是否转hls-fmp4
	EnableMp4     *bool    `json:"enable_mp4,omitempty"`      // 是否mp4录制
	EnableRtsp    *bool    `json:"enable_rtsp,omitempty"`     // 是否转协议为rtsp/webrtc
	EnableRtmp    *bool    `json:"enable_rtmp,omitempty"`     // 是否转协议为rtmp/flv
	EnableTS      *bool    `json:"enable_ts,omitempty"`       // 是否转协议为http-ts/ws-ts
	EnableFmp4    *bool    `json:"enable_fmp4,omitempty"`     // 是否转协议为http-fmp4/ws-fmp4
	EnableAudio   *bool    `json:"enable_audio,omitempty"`    // 转协议是否开启音频
	AddMuteAudio  *bool    `json:"add_mute_audio,omitempty"`  // 转协议无音频时，是否添加静音aac音频
	Mp4SavePath   string   `json:"mp4_save_path,omitempty"`   // mp4录制保存根目录，置空使用默认目录
	Mp4MaxSecond  *int     `json:"mp4_max_second,omitempty"`  // mp4录制切片大小，单位秒
	HlsSavePath   string   `json:"hls_save_path,omitempty"`   // hls保存根目录，置空使用默认目录
	ModifyStamp   *int     `json:"modify_stamp,omitempty"`    // 是否修改原始时间戳，默认值2
	AutoClose     *bool    `json:"auto_close,omitempty"`      // 无人观看时，是否直接关闭
	Latency       *int     `json:"latency,omitempty"`         // srt延时, 单位毫秒
	Passphrase    string   `json:"passphrase,omitempty"`      // srt拉流的密码
}

// AddStreamProxy 添加rtsp/rtmp/hls/srt拉流代理
// 创建一个拉流代理，从指定URL拉取流并在本地提供服务
// 参数:
//   - VHost: 添加的流的虚拟主机，例如__defaultVhost__
//   - App: 添加的流的应用名，例如live
//   - Stream: 添加的流的id名，例如test
//   - URL: 拉流地址，例如rtmp://live.hkstv.hk.lxdns.com/live/hks2
//   - RtpType: rtsp拉流时，拉流方式，0：tcp，1：udp，2：组播
//   - TimeoutSec: 拉流超时时间，单位秒，float类型
//   - RetryCount: 拉流重试次数,不传此参数或传值<=0时，则无限重试
//   - 其他参数: 各种转码和录制选项
//
// 返回: 拉流代理的key，用于后续管理
func (p *ProxyAPI) AddStreamProxy(ctx context.Context, req *AddStreamProxyRequest) (*BaseResponse, error) {
	params := map[string]interface{}{
		"vhost":  req.VHost,
		"app":    req.App,
		"stream": req.Stream,
		"url":    req.URL,
	}

	if req.RtpType != nil {
		params["rtp_type"] = *req.RtpType
	}
	if req.TimeoutSec != nil {
		params["timeout_sec"] = *req.TimeoutSec
	}
	if req.RetryCount != nil {
		params["retry_count"] = *req.RetryCount
	}
	if req.EnableHLS != nil {
		if *req.EnableHLS {
			params["enable_hls"] = "1"
		} else {
			params["enable_hls"] = "0"
		}
	}
	if req.EnableHLSFmp4 != nil {
		if *req.EnableHLSFmp4 {
			params["enable_hls_fmp4"] = "1"
		} else {
			params["enable_hls_fmp4"] = "0"
		}
	}
	if req.EnableMp4 != nil {
		if *req.EnableMp4 {
			params["enable_mp4"] = "1"
		} else {
			params["enable_mp4"] = "0"
		}
	}
	if req.EnableRtsp != nil {
		if *req.EnableRtsp {
			params["enable_rtsp"] = "1"
		} else {
			params["enable_rtsp"] = "0"
		}
	}
	if req.EnableRtmp != nil {
		if *req.EnableRtmp {
			params["enable_rtmp"] = "1"
		} else {
			params["enable_rtmp"] = "0"
		}
	}
	if req.EnableTS != nil {
		if *req.EnableTS {
			params["enable_ts"] = "1"
		} else {
			params["enable_ts"] = "0"
		}
	}
	if req.EnableFmp4 != nil {
		if *req.EnableFmp4 {
			params["enable_fmp4"] = "1"
		} else {
			params["enable_fmp4"] = "0"
		}
	}
	if req.EnableAudio != nil {
		if *req.EnableAudio {
			params["enable_audio"] = "1"
		} else {
			params["enable_audio"] = "0"
		}
	}
	if req.AddMuteAudio != nil {
		if *req.AddMuteAudio {
			params["add_mute_audio"] = "1"
		} else {
			params["add_mute_audio"] = "0"
		}
	}
	if req.Mp4SavePath != "" {
		params["mp4_save_path"] = req.Mp4SavePath
	}
	if req.Mp4MaxSecond != nil {
		params["mp4_max_second"] = *req.Mp4MaxSecond
	}
	if req.HlsSavePath != "" {
		params["hls_save_path"] = req.HlsSavePath
	}
	if req.ModifyStamp != nil {
		params["modify_stamp"] = *req.ModifyStamp
	}
	if req.AutoClose != nil {
		if *req.AutoClose {
			params["auto_close"] = "1"
		} else {
			params["auto_close"] = "0"
		}
	}
	if req.Latency != nil {
		params["latency"] = *req.Latency
	}
	if req.Passphrase != "" {
		params["passphrase"] = req.Passphrase
	}

	respBody, err := p.client.SendRequest(ctx, "GET", "/index/api/addStreamProxy", params)
	if err != nil {
		return nil, fmt.Errorf("添加拉流代理失败: %w", err)
	}

	return ParseResponse(respBody)
}

// DelStreamProxyRequest 关闭拉流代理请求参数
type DelStreamProxyRequest struct {
	Key string `json:"key"` // addStreamProxy接口返回的key
}

// DelStreamProxy 关闭拉流代理
// 关闭指定的拉流代理
// 参数:
//   - Key: addStreamProxy接口返回的key
//
// 返回: 关闭结果
func (p *ProxyAPI) DelStreamProxy(ctx context.Context, req *DelStreamProxyRequest) (*BaseResponse, error) {
	params := map[string]interface{}{
		"key": req.Key,
	}

	respBody, err := p.client.SendRequest(ctx, "GET", "/index/api/delStreamProxy", params)
	if err != nil {
		return nil, fmt.Errorf("关闭拉流代理失败: %w", err)
	}

	return ParseResponse(respBody)
}

// ListStreamProxyRequest 获取拉流代理列表请求参数
type ListStreamProxyRequest struct {
	// 无额外参数，只需要secret
}

// ListStreamProxy 获取拉流代理列表
// 获取所有拉流代理的列表
// 返回: 拉流代理列表信息
func (p *ProxyAPI) ListStreamProxy(ctx context.Context, req *ListStreamProxyRequest) (*BaseResponse, error) {
	respBody, err := p.client.SendRequest(ctx, "GET", "/index/api/listStreamProxy", nil)
	if err != nil {
		return nil, fmt.Errorf("获取拉流代理列表失败: %w", err)
	}

	return ParseResponse(respBody)
}

// AddStreamPusherProxyRequest 添加推流代理请求参数
type AddStreamPusherProxyRequest struct {
	Schema     string   `json:"schema"`                // 推流协议，支持rtsp、rtmp，大小写敏感
	VHost      string   `json:"vhost"`                 // 已注册流的虚拟主机，一般为__defaultVhost__
	App        string   `json:"app"`                   // 已注册流的应用名，例如live
	Stream     string   `json:"stream"`                // 已注册流的id名，例如test
	DstURL     string   `json:"dst_url"`               // 推流地址，需要与schema字段协议一致
	RtpType    *int     `json:"rtp_type,omitempty"`    // rtsp推流时，推流方式，0：tcp，1：udp
	TimeoutSec *float64 `json:"timeout_sec,omitempty"` // 推流超时时间，单位秒，float类型
	RetryCount *int     `json:"retry_count,omitempty"` // 推流重试次数,不传此参数或传值<=0时，则无限重试
}

// AddStreamPusherProxy 添加rtsp/rtmp/srt推流
// 创建一个推流代理，将本地流推送到指定URL
// 参数:
//   - Schema: 推流协议，支持rtsp、rtmp，大小写敏感
//   - VHost: 已注册流的虚拟主机，一般为__defaultVhost__
//   - App: 已注册流的应用名，例如live
//   - Stream: 已注册流的id名，例如test
//   - DstURL: 推流地址，需要与schema字段协议一致
//   - RtpType: rtsp推流时，推流方式，0：tcp，1：udp
//   - TimeoutSec: 推流超时时间，单位秒，float类型
//   - RetryCount: 推流重试次数,不传此参数或传值<=0时，则无限重试
//
// 返回: 推流代理的key，用于后续管理
func (p *ProxyAPI) AddStreamPusherProxy(ctx context.Context, req *AddStreamPusherProxyRequest) (*BaseResponse, error) {
	params := map[string]interface{}{
		"schema":  req.Schema,
		"vhost":   req.VHost,
		"app":     req.App,
		"stream":  req.Stream,
		"dst_url": req.DstURL,
	}

	if req.RtpType != nil {
		params["rtp_type"] = *req.RtpType
	}
	if req.TimeoutSec != nil {
		params["timeout_sec"] = *req.TimeoutSec
	}
	if req.RetryCount != nil {
		params["retry_count"] = *req.RetryCount
	}

	respBody, err := p.client.SendRequest(ctx, "GET", "/index/api/addStreamPusherProxy", params)
	if err != nil {
		return nil, fmt.Errorf("添加推流代理失败: %w", err)
	}

	return ParseResponse(respBody)
}

// ListStreamPusherProxyRequest 获取推流代理列表请求参数
type ListStreamPusherProxyRequest struct {
	// 无额外参数，只需要secret
}

// ListStreamPusherProxy 获取推流代理列表
// 获取所有推流代理的列表
// 返回: 推流代理列表信息
func (p *ProxyAPI) ListStreamPusherProxy(ctx context.Context, req *ListStreamPusherProxyRequest) (*BaseResponse, error) {
	respBody, err := p.client.SendRequest(ctx, "GET", "/index/api/listStreamPusherProxy", nil)
	if err != nil {
		return nil, fmt.Errorf("获取推流代理列表失败: %w", err)
	}

	return ParseResponse(respBody)
}
