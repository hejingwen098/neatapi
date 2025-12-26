// Package auth provides authentication functionality for the NeatLogic API.
// It handles user login, password encryption, and JWT token management.
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

	"github.com/hejingwen098/neatapi/common"
)

// LoginRequest represents the request structure for authentication.
type LoginRequest struct {
	// UserID is the user identifier for authentication.
	UserID string `json:"userid"`
	// Password is the encrypted password for authentication.
	Password string `json:"password"`
}

// LoginResponse represents the response structure for authentication.
type LoginResponse struct {
	// Status indicates the operation status (OK or ERROR).
	Status string `json:"Status"`
	// Message provides additional information about the operation.
	Message string `json:"Message"`
	// JwtToken is the authentication token for subsequent API calls.
	JwtToken string `json:"JwtToken"`
}

// Login performs authentication with the NeatLogic API using default configuration.
// It encrypts the password based on the configuration and returns a JWT token on success.
//
// Returns:
//   - string: The JWT token if authentication is successful
//   - error: An error if authentication fails
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

// LoginWithConfigPath performs authentication with the NeatLogic API using a custom configuration file.
// It initializes the configuration from the specified path, encrypts the password based on the configuration,
// and returns a JWT token on success.
//
// Parameters:
//   - configPath: Path to the configuration file to use
//
// Returns:
//   - string: The JWT token if authentication is successful
//   - error: An error if authentication fails
func LoginWithConfigPath(configPath string) (string, error) {
	common.InitWithConfigPath(configPath)
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
