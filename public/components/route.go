package components

import (
	"github.com/labstack/echo"
	"go-stress-test/components"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var t = &IndexTemplate{
	templates: template.Must(template.ParseGlob("public/views/*.html")),
}
var SupportDataType = []string{
	"int",
	"string",
	"array",
}

func RegisterRoute(e *echo.Echo) {
	e.Renderer = t
	e.GET("/", index)
	e.POST("/add", add)
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

func add(c echo.Context) error {
	registerGlobalConfig(c)
	components.InitRequest()
	components.Start()
	return nil
}

func registerGlobalConfig(c echo.Context) {
	port, _ := strconv.Atoi(c.FormValue("port"))
	remote := components.Remote{
		Protocol: c.FormValue("protocol"),
		Host:     c.FormValue("host"),
		Port:     port,
		Path:     c.FormValue("path"),
		Query:    parseQuery(c),
	}
	method := components.Method{
		Type:        strings.ToLower(c.FormValue("method")),
		ContentType: c.FormValue("content_type"),
	}
	if method.Type == "post" {
		if method.ContentType == "application/json" {
			method.JsonBody = parseJsonBody(c)
		} else {
			method.FormBody = parseFormBody(c)
		}
	}
	threads, _ := strconv.Atoi(c.FormValue("threads"))
	times, _ := strconv.Atoi(c.FormValue("times"))
	test := components.Test{
		Threads: threads,
		Times:   times,
	}
	components.GlobalConfig.Address = &remote
	components.GlobalConfig.Method = &method
	components.GlobalConfig.Test = &test

	components.InitConfig()
}

//解析query参数
func parseQuery(c echo.Context) (query map[string]string) {
	query = make(map[string]string)
	originQuery := c.FormValue("query")
	if originQuery == "" {
		return
	}
	parse := strings.Split(originQuery, "&")
	for _, value := range parse {
		parseValue := strings.Split(value, "=")
		key := parseValue[0]
		val := parseValue[1]
		query[key] = val
	}
	return
}

func parseJsonBody(c echo.Context) (jsonBody map[string]interface{}) {
	jsonBody = make(map[string]interface{})
	if c.FormValue("body") == "" {
		return
	}
	originBody := strings.Split(c.FormValue("body"), "\n")
	for _, v := range originBody {
		parse := strings.Split(v, ":")
		jsonBody[parse[0]] = convertJsonDataType(parse[2], parse[1])
	}
	return
}

func parseFormBody(c echo.Context) (urlValues url.Values) {
	urlValues = make(url.Values)
	if c.FormValue("body") == "" {
		return
	}
	originBody := strings.Split(c.FormValue("body"), "\n")
	for _, v := range originBody {
		parse := strings.Split(v, ":")
		urlValues[parse[0]] = convertFormDataType(parse[2], parse[1])
	}
	return
}

func convertJsonDataType(data string, dataType string) (afterConvert interface{}) {
	afterConvert = data
	support := false
	for _, supportType := range SupportDataType {
		if dataType == supportType {
			support = true
		}
	}
	if support == false {
		return
	}
	switch dataType {
	case "int":
		afterConvert, _ = strconv.Atoi(data)
		break
	case "string":
		afterConvert = string(data)
		break
	case "array":
		afterConvert = strings.Split(data, ",")
		break
	default:
		break
	}
	return
}

func convertFormDataType(data string, dataType string) (afterConvert []string) {
	support := false
	for _, supportType := range SupportDataType {
		if dataType == supportType {
			support = true
		}
	}
	if support == false {
		return
	}
	switch dataType {
	case "int":
	case "string":
		afterConvert = []string{data}
		break
	case "array":
		afterConvert = strings.Split(data, ",")
		break
	default:
		break
	}
	return
}
