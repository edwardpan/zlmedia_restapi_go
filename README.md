# zlmediaKit Go SDK

这是一个用于与zlmediaKit媒体服务器进行API交互的Go语言工具库。

## 功能特性

- 支持zlmediaKit所有RESTful API接口
- 自动处理secret参数认证
- 支持GET和POST请求方式
- 统一的响应结构处理
- 详细的接口注释和参数说明
- 模块化设计，按功能分类

## 安装

```bash
go get github.com/edwardpan/zlmedia_restapi_go
```

## 快速开始

### 1. 初始化客户端

```go
package main

import (
    "context"
    "github.com/edwardpan/zlmedia_restapi_go"
)

func main() {
    // 配置客户端
    config := &zlmedia_restapi_go.Config{
        BaseURL: "http://127.0.0.1:80",
        Secret:  "your-secret-key",
        Timeout: 30,
    }
    
    // 初始化客户端
    client := zlmedia_restapi_go.InitClient(config)
    ctx := context.Background()
}
```

### 2. 使用API

```go
// 获取服务器API列表
serverAPI := zlmedia_restapi_go.GetServerAPI()
resp, err := serverAPI.GetApiList(ctx, &zlmedia_restapi_go.GetApiListRequest{})
if err != nil {
    log.Printf("错误: %v", err)
    return
}

// 检查响应
if resp.Code == 0 {
    fmt.Println("请求成功")
    fmt.Printf("数据: %+v\n", resp.Data)
} else {
    fmt.Printf("请求失败: %s\n", resp.Msg)
}
```

## API模块

### 1. 服务器管理 (ServerAPI)

```go
serverAPI := zlmedia_restapi_go.GetServerAPI()

// 获取API列表
resp, err := serverAPI.GetApiList(ctx, &zlmedia_restapi_go.GetApiListRequest{})

// 获取服务器配置
resp, err := serverAPI.GetServerConfig(ctx, &zlmedia_restapi_go.GetServerConfigRequest{})

// 设置服务器配置
resp, err := serverAPI.SetServerConfig(ctx, &zlmedia_restapi_go.SetServerConfigRequest{
    Key:   "api.secret",
    Value: "new-secret",
})

// 重启服务器
resp, err := serverAPI.RestartServer(ctx, &zlmedia_restapi_go.RestartServerRequest{})
```

### 2. 流媒体管理 (MediaAPI)

```go
mediaAPI := zlmedia_restapi_go.GetMediaAPI()

// 获取流列表
resp, err := mediaAPI.GetMediaList(ctx, &zlmedia_restapi_go.GetMediaListRequest{
    Schema: "rtmp",
    VHost:  "__defaultVhost__",
    App:    "live",
})

// 关闭单个流
resp, err := mediaAPI.CloseStream(ctx, &zlmedia_restapi_go.CloseStreamRequest{
    Schema: "rtmp",
    VHost:  "__defaultVhost__",
    App:    "live",
    Stream: "test",
    Force:  true,
})

// 批量关闭流
resp, err := mediaAPI.CloseStreams(ctx, &zlmedia_restapi_go.CloseStreamsRequest{
    Schema: "rtmp",
    VHost:  "__defaultVhost__",
    App:    "live",
    Force:  true,
})
```

### 3. 代理管理 (ProxyAPI)

```go
proxyAPI := zlmedia_restapi_go.GetProxyAPI()

// 添加拉流代理
enableHLS := true
resp, err := proxyAPI.AddStreamProxy(ctx, &zlmedia_restapi_go.AddStreamProxyRequest{
    VHost:     "__defaultVhost__",
    App:       "live",
    Stream:    "test",
    URL:       "rtmp://example.com/live/stream",
    EnableHLS: &enableHLS,
})

// 获取拉流代理列表
resp, err := proxyAPI.ListStreamProxy(ctx, &zlmedia_restapi_go.ListStreamProxyRequest{})

// 删除拉流代理
resp, err := proxyAPI.DelStreamProxy(ctx, &zlmedia_restapi_go.DelStreamProxyRequest{
    Key: "proxy-key",
})

// 添加推流代理
resp, err := proxyAPI.AddStreamPusherProxy(ctx, &zlmedia_restapi_go.AddStreamPusherProxyRequest{
    Schema: "rtmp",
    VHost:  "__defaultVhost__",
    App:    "live",
    Stream: "test",
    DstURL: "rtmp://push.example.com/live/stream",
})
```

### 4. 录制管理 (RecordAPI)

```go
recordAPI := zlmedia_restapi_go.GetRecordAPI()

// 开始录制
resp, err := recordAPI.StartRecord(ctx, &zlmedia_restapi_go.StartRecordRequest{
    Type:   1, // 1为mp4，0为hls
    VHost:  "__defaultVhost__",
    App:    "live",
    Stream: "test",
})

// 停止录制
resp, err := recordAPI.StopRecord(ctx, &zlmedia_restapi_go.StopRecordRequest{
    Type:   1,
    VHost:  "__defaultVhost__",
    App:    "live",
    Stream: "test",
})

// 检查录制状态
resp, err := recordAPI.IsRecording(ctx, &zlmedia_restapi_go.IsRecordingRequest{
    Type:   1,
    VHost:  "__defaultVhost__",
    App:    "live",
    Stream: "test",
})

// 获取录制文件列表
resp, err := recordAPI.GetMp4RecordFile(ctx, &zlmedia_restapi_go.GetMp4RecordFileRequest{
    VHost:  "__defaultVhost__",
    App:    "live",
    Stream: "test",
    Period: "2024-01-01",
})
```

### 5. RTP管理 (RTPAPI)

```go
rtpAPI := zlmedia_restapi_go.GetRTPAPI()

// 创建RTP接收端口
enableTcp := 0
resp, err := rtpAPI.OpenRtpServer(ctx, &zlmedia_restapi_go.OpenRtpServerRequest{
    Port:      10000,
    EnableTcp: &enableTcp,
    StreamID:  "gb28181_stream",
})

// 关闭RTP接收端口
resp, err := rtpAPI.CloseRtpServer(ctx, &zlmedia_restapi_go.CloseRtpServerRequest{
    StreamID: "gb28181_stream",
})

// 启动RTP推流
resp, err := rtpAPI.StartSendRtp(ctx, &zlmedia_restapi_go.StartSendRtpRequest{
    VHost:   "__defaultVhost__",
    App:     "live",
    Stream:  "test",
    Ssrc:    "12345678",
    DstURL:  "192.168.1.100",
    DstPort: 10000,
})

// 停止RTP推流
resp, err := rtpAPI.StopSendRtp(ctx, &zlmedia_restapi_go.StopSendRtpRequest{
    VHost:  "__defaultVhost__",
    App:    "live",
    Stream: "test",
    Ssrc:   "12345678",
})
```

### 6. 会话管理 (SessionAPI)

```go
sessionAPI := zlmedia_restapi_go.GetSessionAPI()

// 获取会话列表
resp, err := sessionAPI.ListSession(ctx, &zlmedia_restapi_go.ListSessionRequest{})

// 断开单个连接
resp, err := sessionAPI.KickSession(ctx, &zlmedia_restapi_go.KickSessionRequest{
    ID: "session-id",
})

// 批量断开连接
resp, err := sessionAPI.KickSessions(ctx, &zlmedia_restapi_go.KickSessionsRequest{
    Schema:    "rtmp",
    VHost:     "__defaultVhost__",
    App:       "live",
    Stream:    "test",
    LocalPort: 1935,
})
```

### 7. WebRTC管理 (WebRTCAPI)

```go
webrtcAPI := zlmedia_restapi_go.GetWebRTCAPI()

// WebRTC offer/answer交换
resp, err := webrtcAPI.WebRTC(ctx, &zlmedia_restapi_go.WebRTCRequest{
    Api:    "play",
    Type:   "offer",
    SDP:    "sdp-content",
    VHost:  "__defaultVhost__",
    App:    "live",
    Stream: "test",
})
```

## 响应结构

所有API调用都返回统一的响应结构：

```go
type BaseResponse struct {
    Code int                    `json:"code"` // 错误代码，0表示成功
    Msg  string                 `json:"msg"`  // 错误信息
    Data map[string]interface{} `json:"data,omitempty"` // 响应数据
}
```

## 错误处理

```go
resp, err := serverAPI.GetApiList(ctx, &zlmedia_restapi_go.GetApiListRequest{})
if err != nil {
    // 网络错误或其他系统错误
    log.Printf("请求失败: %v", err)
    return
}

if resp.Code != 0 {
    // zlmedia_restapi_goKit返回的业务错误
    log.Printf("业务错误: %s (code: %d)", resp.Msg, resp.Code)
    return
}

// 请求成功，处理数据
fmt.Printf("数据: %+v\n", resp.Data)
```

## 配置选项

```go
type Config struct {
    BaseURL string // zlmedia_restapi_goKit服务器地址，如: http://127.0.0.1:80
    Secret  string // API密钥
    Timeout int    // 请求超时时间（秒），默认30秒
}
```

## 注意事项

1. 所有API调用都会自动添加`secret`参数进行认证
2. GET请求参数通过URL查询字符串传递
3. POST请求参数通过JSON格式的请求体传递
4. 布尔类型参数在传递时会转换为字符串"0"或"1"
5. 可选参数使用指针类型，nil值表示不传递该参数

## 参考文档

- [zlmediaKit官方文档](https://docs.zlmedia.com/)
- [zlmediaKit RESTful API文档](https://docs.zlmedia.com/zh/guide/media_server/restful_api.html)