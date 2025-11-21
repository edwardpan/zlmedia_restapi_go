package zlmedia

import (
	"context"
	"fmt"
)

// SessionAPI 会话管理相关API
type SessionAPI struct {
	client *Client
}

// NewSessionAPI 创建会话API实例
func NewSessionAPI(client *Client) *SessionAPI {
	return &SessionAPI{client: client}
}

// GetSessionAPI 获取会话API实例
func GetSessionAPI() *SessionAPI {
	return NewSessionAPI(GetClient())
}

// GetAllSessionRequest 获取Session列表请求参数
type GetAllSessionRequest struct {
	LocalPort *int   `json:"local_port,omitempty"` // 筛选本机端口，例如筛选rtsp链接：554
	PeerIP    string `json:"peer_ip,omitempty"`    // 筛选客户端ip
}

// GetAllSession 获取Session列表
// 获取ZLMediaKit中所有TCP连接会话的列表
// 参数:
//   - LocalPort: 筛选本机端口，例如筛选rtsp链接：554
//   - PeerIP: 筛选客户端ip
//
// 返回: Session列表信息
func (s *SessionAPI) GetAllSession(ctx context.Context, req *GetAllSessionRequest) (*BaseResponse, error) {
	params := make(map[string]interface{})
	if req.LocalPort != nil {
		params["local_port"] = *req.LocalPort
	}
	if req.PeerIP != "" {
		params["peer_ip"] = req.PeerIP
	}

	respBody, err := s.client.SendRequest(ctx, "GET", "/index/api/getAllSession", params)
	if err != nil {
		return nil, fmt.Errorf("获取Session列表失败: %w", err)
	}

	return ParseResponse(respBody)
}

// KickSessionRequest 断开tcp连接请求参数
type KickSessionRequest struct {
	ID string `json:"id"` // 客户端唯一id，可以通过getAllSession接口获取
}

// KickSession 断开tcp连接
// 断开指定的TCP连接会话
// 参数:
//   - ID: 客户端唯一id，可以通过getAllSession接口获取
//
// 返回: 断开连接结果
func (s *SessionAPI) KickSession(ctx context.Context, req *KickSessionRequest) (*BaseResponse, error) {
	params := map[string]interface{}{
		"id": req.ID,
	}

	respBody, err := s.client.SendRequest(ctx, "GET", "/index/api/kick_session", params)
	if err != nil {
		return nil, fmt.Errorf("断开连接失败: %w", err)
	}

	return ParseResponse(respBody)
}

// KickSessionsRequest 批量断开tcp连接请求参数
type KickSessionsRequest struct {
	LocalPort *int   `json:"local_port,omitempty"` // 筛选本机端口，例如筛选rtsp链接：554
	PeerIP    string `json:"peer_ip,omitempty"`    // 筛选客户端ip
}

// KickSessions 批量断开tcp连接
// 批量断开符合条件的TCP连接会话
// 参数:
//   - LocalPort: 筛选本机端口，例如筛选rtsp链接：554
//   - PeerIP: 筛选客户端ip
//
// 返回: 批量断开连接结果
func (s *SessionAPI) KickSessions(ctx context.Context, req *KickSessionsRequest) (*BaseResponse, error) {
	params := make(map[string]interface{})
	if req.LocalPort != nil {
		params["local_port"] = *req.LocalPort
	}
	if req.PeerIP != "" {
		params["peer_ip"] = req.PeerIP
	}

	respBody, err := s.client.SendRequest(ctx, "GET", "/index/api/kick_sessions", params)
	if err != nil {
		return nil, fmt.Errorf("批量断开连接失败: %w", err)
	}

	return ParseResponse(respBody)
}
