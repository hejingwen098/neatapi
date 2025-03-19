package common

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	Config       Configs // 假设您已经定义了ConfigStruct来匹配您的配置文件
	NeatlogicUri string
)

type Auth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Encrypt  string `yaml:"encrypt"`
}
type Neatlogic struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Tenant string `yaml:"tenant"`
}
type Global struct {
	Auth      Auth      `yaml:"auth"`
	Neatlogic Neatlogic `yaml:"neatlogic"`
}
type Configs struct {
	Global Global `yaml:"global"`
}

func init() {
	// 初始化配置文件
	Config = Configs{}
	configFile, err := os.Open("./config.yml")
	if err != nil {
		fmt.Printf("Error opening config file: %s", err)
		// return nil, err
	}
	defer configFile.Close()

	// 解析配置文件
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Printf("Error parsing config file: %s", err)
		// return nil, err
	}
	NeatlogicUri = fmt.Sprintf("http://%s:%d/%s", Config.Global.Neatlogic.Host, Config.Global.Neatlogic.Port, Config.Global.Neatlogic.Tenant)

}
