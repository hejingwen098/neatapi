package auth

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"neatapi/interval/common"
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

func Login() (string, error) {
	api_login := fmt.Sprintf("%s/login/check", common.NeatlogicUri)
	var JwtToken, encryptedPass string
	// Password encryption
	if common.Config.Global.Auth.Encrypt == "base64" {
		encryptedPass = "{BS}" + base64.StdEncoding.EncodeToString([]byte(common.Config.Global.Auth.Password))
	} else {
		// Default to MD5
		hasher := md5.New()
		hasher.Write([]byte(common.Config.Global.Auth.Password))
		encryptedPass = "{MD5}" + hex.EncodeToString(hasher.Sum(nil))
	}

	// Create request body
	reqBody := LoginRequest{
		UserID:   common.Config.Global.Auth.Username,
		Password: encryptedPass,
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", errors.New("failed to marshal login request")
	}

	// Make HTTP request
	resp, err := http.Post(api_login, "application/json", strings.NewReader(string(jsonData)))
	if err != nil {
		return "", errors.New("failed to login")
	}
	defer resp.Body.Close()

	// Parse response
	var loginResp LoginResponse
	err = json.NewDecoder(resp.Body).Decode(&loginResp)
	if err != nil {
		return "", errors.New("failed to parse login response")
	}

	// Handle response
	if loginResp.Status == "OK" {
		// Store JWT token
		if loginResp.JwtToken != "" {
			JwtToken = loginResp.JwtToken
		}
		return JwtToken, nil
	} else if loginResp.Status == "ERROR" {
		return "", errors.New(loginResp.Message)
	}

	return "", errors.New("unknown authentication error")
}
