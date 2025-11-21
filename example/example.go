package main

import (
	"context"
	"fmt"
	"log"

	zlmedia "github.com/edwardpan/zlmedia_restapi_go"
)

func main() {
	// 初始化ZLMediaKit客户端
	config := zlmedia.Config{
		BaseURL: "http://127.0.0.1:80",
		Secret:  "035c73f7-bb6b-4889-a715-d9eb2d1925cc",
		Timeout: 30,
	}

	_ = zlmedia.InitClient(config)
	ctx := context.Background()

	// 示例1: 获取API列表
	fmt.Println("=== 获取API列表 ===")
	serverAPI := zlmedia.GetServerAPI()
	apiListResp, err := serverAPI.GetApiList(ctx, &zlmedia.GetApiListRequest{})
	if err != nil {
		log.Printf("获取API列表失败: %v", err)
	} else {
		fmt.Printf("API列表响应: %+v\n", apiListResp)
	}

	// 示例2: 获取流列表
	fmt.Println("\n=== 获取流列表 ===")
	mediaAPI := zlmedia.GetMediaAPI()
	mediaListResp, err := mediaAPI.GetMediaList(ctx, &zlmedia.GetMediaListRequest{
		Schema: "rtmp",
		VHost:  "__defaultVhost__",
		App:    "live",
	})
	if err != nil {
		log.Printf("获取流列表失败: %v", err)
	} else {
		fmt.Printf("流列表响应: %+v\n", mediaListResp)
	}

	// 示例3: 添加拉流代理
	fmt.Println("\n=== 添加拉流代理 ===")
	proxyAPI := zlmedia.GetProxyAPI()
	enableHLS := true
	enableMp4 := false
	addProxyResp, err := proxyAPI.AddStreamProxy(ctx, &zlmedia.AddStreamProxyRequest{
		VHost:     "__defaultVhost__",
		App:       "live",
		Stream:    "test",
		URL:       "rtmp://live.hkstv.hk.lxdns.com/live/hks2",
		EnableHLS: &enableHLS,
		EnableMp4: &enableMp4,
	})
	if err != nil {
		log.Printf("添加拉流代理失败: %v", err)
	} else {
		fmt.Printf("添加拉流代理响应: %+v\n", addProxyResp)
	}

	// 示例4: 开始录制
	fmt.Println("\n=== 开始录制 ===")
	recordAPI := zlmedia.GetRecordAPI()
	startRecordResp, err := recordAPI.StartRecord(ctx, &zlmedia.StartRecordRequest{
		Type:   1, // mp4录制
		VHost:  "__defaultVhost__",
		App:    "live",
		Stream: "test",
	})
	if err != nil {
		log.Printf("开始录制失败: %v", err)
	} else {
		fmt.Printf("开始录制响应: %+v\n", startRecordResp)
	}

	// 示例5: 创建RTP接收端口
	fmt.Println("\n=== 创建RTP接收端口 ===")
	rtpAPI := zlmedia.GetRTPAPI()
	enableTcp := 0
	openRtpResp, err := rtpAPI.OpenRtpServer(ctx, &zlmedia.OpenRtpServerRequest{
		Port:      0, // 随机端口
		EnableTcp: &enableTcp,
		StreamID:  "gb28181_test",
	})
	if err != nil {
		log.Printf("创建RTP接收端口失败: %v", err)
	} else {
		fmt.Printf("创建RTP接收端口响应: %+v\n", openRtpResp)
	}

	// 示例6: 获取会话列表
	fmt.Println("\n=== 获取会话列表 ===")
	sessionAPI := zlmedia.GetSessionAPI()
	sessionListResp, err := sessionAPI.GetAllSession(ctx, &zlmedia.GetAllSessionRequest{})
	if err != nil {
		log.Printf("获取会话列表失败: %v", err)
	} else {
		fmt.Printf("会话列表响应: %+v\n", sessionListResp)
	}

	// 示例7: 获取服务器配置
	fmt.Println("\n=== 获取服务器配置 ===")
	configResp, err := serverAPI.GetServerConfig(ctx, &zlmedia.GetServerConfigRequest{})
	if err != nil {
		log.Printf("获取服务器配置失败: %v", err)
	} else {
		fmt.Printf("服务器配置响应: %+v\n", configResp)
	}

	fmt.Println("\n=== 示例完成 ===")
}
