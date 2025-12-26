// Package common provides common utilities and configuration management for the NeatAPI SDK.
// It handles configuration loading, parsing, and global variable initialization.
package common

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	// Config holds the application configuration loaded from config.yml.
	Config Configs
	// NeatlogicUri is the base URL for the NeatLogic API, constructed from configuration.
	NeatlogicUri string
)

// Auth represents the authentication configuration section.
type Auth struct {
	// Username is the user identifier for authentication.
	Username string `yaml:"username"`
	// Password is the password for authentication.
	Password string `yaml:"password"`
	// Encrypt specifies the encryption method (base64 or md5).
	Encrypt string `yaml:"encrypt"`
}

// Neatlogic represents the NeatLogic service configuration section.
type Neatlogic struct {
	// Host is the hostname or IP address of the NeatLogic server.
	Host string `yaml:"host"`
	// Port is the port number of the NeatLogic server.
	Port int `yaml:"port"`
	// Tenant is the tenant identifier for the NeatLogic instance.
	Tenant string `yaml:"tenant"`
}

// Global represents the global configuration section.
type Global struct {
	// Auth contains authentication configuration.
	Auth Auth `yaml:"auth"`
	// Neatlogic contains NeatLogic service configuration.
	Neatlogic Neatlogic `yaml:"neatlogic"`
}

// Configs represents the complete configuration structure.
type Configs struct {
	// Global contains global configuration settings.
	Global Global `yaml:"global"`
}

// init initializes the configuration when the package is loaded.
// It reads the default config.yml file and parses the configuration.
func init() {
	// Initialize configuration file
	Config = Configs{}
	configFile, err := os.Open("./config.yml")
	if err != nil {
		fmt.Printf("Error opening config file: %s", err)
		// return nil, err
	}
	defer configFile.Close()

	// Parse configuration file
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Printf("Error parsing config file: %s", err)
		// return nil, err
	}
	NeatlogicUri = fmt.Sprintf("http://%s:%d/%s", Config.Global.Neatlogic.Host, Config.Global.Neatlogic.Port, Config.Global.Neatlogic.Tenant)

}

// InitWithConfigPath initializes the configuration from a custom configuration file path.
// It reads the specified configuration file and parses it into the global Config variable.
//
// Parameters:
//   - configPath: Path to the configuration file to use
func InitWithConfigPath(configPath string) {
	// Initialize configuration file
	Config = Configs{}
	configFile, err := os.Open(configPath)
	if err != nil {
		fmt.Printf("Error opening config file: %s", err)
		// return nil, err
	}
	defer configFile.Close()

	// Parse configuration file
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Printf("Error parsing config file: %s", err)
		// return nil, err
	}
	NeatlogicUri = fmt.Sprintf("http://%s:%d/%s", Config.Global.Neatlogic.Host, Config.Global.Neatlogic.Port, Config.Global.Neatlogic.Tenant)

}
