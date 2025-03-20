package neatlogic

import (
	"fmt"
	"io"
	"net/http"
)

type LoginRequest struct {
	UserID   string `json:"userid"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status   string `json:"Status"`
	Message  string `json:"Message"`
	JwtToken string `json:"JwtToken"`
}

func SendRequest(req *http.Request, JwtToken string) ([]byte, error) {
	client := &http.Client{}
	// 设置JWT认证头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+JwtToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析响应
	respBody, err := ParseResourceResponse(resp)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func ParseResourceResponse(resp *http.Response) ([]byte, error) {
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
