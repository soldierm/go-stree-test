package components

import (
	"path/filepath"
	"strings"
)

const (
	DEFAULT_THREADS    = 20     /* 默认线程数 */
	DEFAULT_METHOD     = "GET"  /* 默认GET方法 */
	DEFAULT_PROTOCOL   = "http" /* 默认http请求 */
	DEFAULT_HTTP_PORT  = 80     /* http默认80端口 */
	DEFAULT_HTTPS_PORT = 443    /* https默认443端口 */
)

type Config struct {
	Address Remote `yaml:"api_address"`
	Method  string `yaml:"method"`
	Test    Test   `yaml:"test"`
}

type Remote struct {
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Path     string `yaml:"path"`
}

type Test struct {
	Threads int `yaml:"threads"`
	Times   int `yaml:"times"`
}

var GlobalConfig Config

func init() {
	parseYml()
	validate()
}

//解析yaml文件，读取配置
func parseYml() {
	configFile, _ := filepath.Abs("./config/config.yml")
	GlobalConfig, _ = ParseYamlFromFile(configFile)
}

//验证配置信息
func validate() {
	if strings.ToUpper(GlobalConfig.Method) != DEFAULT_METHOD {
		GlobalConfig.Method = DEFAULT_METHOD
	}
	if GlobalConfig.Address.Protocol == DEFAULT_PROTOCOL {
		GlobalConfig.Address.Port = DEFAULT_HTTP_PORT
	} else {
		GlobalConfig.Address.Port = DEFAULT_HTTPS_PORT
	}
	if GlobalConfig.Test.Threads <= 0 {
		GlobalConfig.Test.Threads = DEFAULT_THREADS
	}
}

//url转成string
func (remote *Remote) UriToString() string {
	port := string(remote.Port)
	if remote.Protocol == DEFAULT_PROTOCOL && remote.Port == DEFAULT_HTTP_PORT {
		port = ""
	} else if remote.Protocol == "https" && remote.Port == DEFAULT_HTTPS_PORT {
		port = ""
	}
	return remote.Protocol + "://" + remote.Host + port + remote.Path
}
