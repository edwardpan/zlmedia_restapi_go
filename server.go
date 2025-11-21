package zlmedia

import (
	"context"
	"fmt"
)

// ServerAPI 服务器管理相关API
type ServerAPI struct {
	client *Client
}

// NewServerAPI 创建服务器API实例
func NewServerAPI(client *Client) *ServerAPI {
	return &ServerAPI{client: client}
}

// GetServerAPI 获取服务器API实例
func GetServerAPI() *ServerAPI {
	return NewServerAPI(GetClient())
}

// GetApiListRequest 获取服务器api列表请求参数
type GetApiListRequest struct {
	// 无额外参数，只需要secret
}

// GetApiList 获取服务器api列表
// 获取ZLMediaKit支持的所有API接口列表
// 返回: API接口列表信息
func (s *ServerAPI) GetApiList(ctx context.Context, req *GetApiListRequest) (*BaseResponse, error) {
	respBody, err := s.client.SendRequest(ctx, "GET", "/index/api/getApiList", nil)
	if err != nil {
		return nil, fmt.Errorf("获取API列表失败: %w", err)
	}

	return ParseResponse(respBody)
}

// GetThreadsLoadRequest 获取网络线程负载请求参数
type GetThreadsLoadRequest struct {
	// 无额外参数，只需要secret
}

// GetThreadsLoad 获取网络线程负载
// 获取ZLMediaKit网络线程的负载情况
// 返回: 网络线程负载信息
func (s *ServerAPI) GetThreadsLoad(ctx context.Context, req *GetThreadsLoadRequest) (*BaseResponse, error) {
	respBody, err := s.client.SendRequest(ctx, "GET", "/index/api/getThreadsLoad", nil)
	if err != nil {
		return nil, fmt.Errorf("获取网络线程负载失败: %w", err)
	}

	return ParseResponse(respBody)
}

// GetStatisticRequest 获取主要对象个数请求参数
type GetStatisticRequest struct {
	// 无额外参数，只需要secret
}

// GetStatistic 获取主要对象个数
// 获取ZLMediaKit中主要对象的统计信息，如流的数量等
// 返回: 主要对象统计信息
func (s *ServerAPI) GetStatistic(ctx context.Context, req *GetStatisticRequest) (*BaseResponse, error) {
	respBody, err := s.client.SendRequest(ctx, "GET", "/index/api/getStatistic", nil)
	if err != nil {
		return nil, fmt.Errorf("获取统计信息失败: %w", err)
	}

	return ParseResponse(respBody)
}

// GetWorkThreadsLoadRequest 获取后台线程负载请求参数
type GetWorkThreadsLoadRequest struct {
	// 无额外参数，只需要secret
}

// GetWorkThreadsLoad 获取后台线程负载
// 获取ZLMediaKit后台工作线程的负载情况
// 返回: 后台线程负载信息
func (s *ServerAPI) GetWorkThreadsLoad(ctx context.Context, req *GetWorkThreadsLoadRequest) (*BaseResponse, error) {
	respBody, err := s.client.SendRequest(ctx, "GET", "/index/api/getWorkThreadsLoad", nil)
	if err != nil {
		return nil, fmt.Errorf("获取后台线程负载失败: %w", err)
	}

	return ParseResponse(respBody)
}

// GetServerConfigRequest 获取服务器配置请求参数
type GetServerConfigRequest struct {
	// 无额外参数，只需要secret
}

// GetServerConfig 获取服务器配置
// 获取ZLMediaKit的完整配置信息
// 返回: 服务器配置信息
func (s *ServerAPI) GetServerConfig(ctx context.Context, req *GetServerConfigRequest) (*BaseResponse, error) {
	respBody, err := s.client.SendRequest(ctx, "GET", "/index/api/getServerConfig", nil)
	if err != nil {
		return nil, fmt.Errorf("获取服务器配置失败: %w", err)
	}

	return ParseResponse(respBody)
}

// SetServerConfigRequest 设置服务器配置请求参数
type SetServerConfigRequest struct {
	// 配置项，格式为 "section.key": "value"
	// 例如: "api.apiDebug": "0"
	Config map[string]string `json:"config"`
}

// SetServerConfig 设置服务器配置
// 动态修改ZLMediaKit的配置项
// 参数:
//   - Config: 配置项映射，键为"section.key"格式，值为配置值
//
// 返回: 设置结果
func (s *ServerAPI) SetServerConfig(ctx context.Context, req *SetServerConfigRequest) (*BaseResponse, error) {
	params := make(map[string]interface{})
	for key, value := range req.Config {
		params[key] = value
	}

	respBody, err := s.client.SendRequest(ctx, "GET", "/index/api/setServerConfig", params)
	if err != nil {
		return nil, fmt.Errorf("设置服务器配置失败: %w", err)
	}

	return ParseResponse(respBody)
}

// RestartServerRequest 重启服务器请求参数
type RestartServerRequest struct {
	// 无额外参数，只需要secret
}

// RestartServer 重启服务器
// 重启ZLMediaKit服务器，注意这会中断所有正在进行的流
// 返回: 重启结果
func (s *ServerAPI) RestartServer(ctx context.Context, req *RestartServerRequest) (*BaseResponse, error) {
	respBody, err := s.client.SendRequest(ctx, "GET", "/index/api/restartServer", nil)
	if err != nil {
		return nil, fmt.Errorf("重启服务器失败: %w", err)
	}

	return ParseResponse(respBody)
}
