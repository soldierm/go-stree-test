package components

import (
	"log"
	url2 "net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	DEFAULT_THREADS    = 20     /* 默认线程数 */
	DEFAULT_PROTOCOL   = "http" /* 默认http请求 */
	DEFAULT_HTTP_PORT  = 80     /* http默认80端口 */
	DEFAULT_HTTPS_PORT = 443    /* https默认443端口 */
)

type Config struct {
	Address *Remote `yaml:"api_address"`
	Method  *Method `yaml:"method"`
	Test    *Test   `yaml:"test"`
}

type Remote struct {
	Protocol string            `yaml:"protocol"`
	Host     string            `yaml:"host"`
	Port     int               `yaml:"port"`
	Path     string            `yaml:"path"`
	Query    map[string]string `yaml:"query"`
}

type Method struct {
	Type        string                 `yaml:"type"`
	FormBody    url2.Values            `yaml:"form_body"`
	JsonBody    map[string]interface{} `yaml:"json_body"`
	ContentType string                 `yaml:"content_type"`
}

type Test struct {
	Threads int `yaml:"threads"`
	Times   int `yaml:"times"`
}

var GlobalConfig Config
var SupportMethods = []string{
	"get",
	"post",
}

//初始化操作
func InitConfig() {
	if isTerminal() {
		parseYml()
		validate()
	}
}

//解析yaml文件，读取配置
func parseYml() {
	configFile, _ := filepath.Abs("./config/config.yml")
	GlobalConfig, _ = ParseYamlFromFile(configFile)
}

//验证配置信息
func validate() {
	method := strings.ToLower(GlobalConfig.Method.Type)
	support := false
	for _, supportMethod := range SupportMethods {
		if supportMethod == method {
			support = true
		}
	}
	if support == false {
		log.Fatalf("不支持的请求方式：%s", method)
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
	query := ""
	if len(remote.Query) > 0 {
		query = "?"
		for attr, value := range remote.Query {
			query += attr + "=" + value + "&"
		}
		query = strings.Trim(query, "&")
	}
	return remote.Protocol + "://" + remote.Host + port + remote.Path + query
}

func isTerminal() bool {
	return os.Getenv("CURRENT_ENV") == "terminal"
}
