package config

import (
	"errors"
	"fmt"
	"github.com/wjpxxx/letgo/encry"
	"github.com/wjpxxx/letgo/httpclient"
	"github.com/wjpxxx/letgo/lib"
	"sort"
	"strings"
)

//Config
type Config struct {
	BaseURL     string `json:"baseURL"`
	AppKey      string `json:"app_key"`
	AppSecret   string `json:"app_secret"`
	RedirectURL string `json:"redirect_url"`
	AccessToken string `json:"access_token"`
	ShopId      string `json:"shop_id"`
}

//String
func (c *Config) String() string {
	return lib.ObjectToString(c)
}

//GetApiURL
func (c *Config) GetApiURL(apiPath string) string {
	return fmt.Sprintf("%s%s", c.BaseURL, apiPath)
}

//GetCommonParam
func (c *Config) GetCommonParam(method string) lib.InRow {
	ti := lib.Time()
	param := lib.InRow{
		"app_key":      c.AppKey,
		"timestamp":    ti,
		"access_token": c.AccessToken,
		"shop_id":      c.ShopId,
	}
	return param
}

//HttpGet
func (c *Config) HttpGet(method string, data interface{}, out interface{}) error {
	return c.Http("GET", method, data, out)
}

//HttpPost
func (c *Config) HttpPost(method string, data interface{}, out interface{}) error {
	return c.Http("POST", method, data, out)
}

//HttpPostFile
func (c *Config) HttpPostFile(method string, data interface{}, out interface{}) error {
	return c.Http("POSTFILE", method, data, out)
}

//Http 请求
func (c *Config) Http(requestMethod, method string, data interface{}, out interface{}) error {
	param := c.GetCommonParam(method)
	inputParam:=data.(lib.InRow)
	allParam:=lib.MergeInRow(param,inputParam)
	param["sign"]=Sign(c.AppSecret,method,allParam)
	apiUrl:=c.BaseURL
	fullURL := fmt.Sprintf("%s%s?%s", apiUrl, method,httpclient.HttpBuildQuery(param))
	
	ihttp := httpclient.New().WithTimeOut(120)
	var result *httpclient.HttpResponse
	if requestMethod == "GET" {
		result = ihttp.Get(fullURL, inputParam)
	} else {
		result = ihttp.Post(fullURL, inputParam)
	}
	//fmt.Println(result.Dump)
	if result.Err != "" {
		return errors.New(result.Err)
	}
	if result.Code != 200 {
		return errors.New("请求失败")
	}
	lib.StringToObject(result.Body(), out)
	return nil
}

//New
func New(apiURL, AppKey, AppSecret, AccessToken, ShopId string, redirectURL string) *Config {
	return &Config{
		apiURL,
		AppKey,
		AppSecret,
		redirectURL,
		AccessToken,
		ShopId,
	}
}

//Sign
func Sign(appSecret, method string, param lib.InRow) string {
	tmpParam := lib.InRow{}
	//排除掉文件上传
	for k, v := range param {
		if k[0] != '@' {
			tmpParam[k] = v
		}
	}
	query := ""
	sortedKeys := make([]string, 0)
	for k, _ := range tmpParam {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	for _, k := range sortedKeys {
		if vs, ok := tmpParam[k].(string); ok {
			query += fmt.Sprintf("%s%s", k, vs)
		} else if vs, ok := tmpParam[k].(int); ok {
			query += fmt.Sprintf("%s%d", k, vs)
		} else if vs, ok := tmpParam[k].(int64); ok {
			query += fmt.Sprintf("%s%d", k, vs)
		} else if vs, ok := tmpParam[k].(int32); ok {
			query += fmt.Sprintf("%s%d", k, vs)
		} else if vs, ok := tmpParam[k].(float32); ok {
			query += fmt.Sprintf("%s%f", k, vs)
		} else if vs, ok := tmpParam[k].(float64); ok {
			query += fmt.Sprintf("%s%f", k, vs)
		} else if vs, ok := tmpParam[k].(bool); ok {
			query += fmt.Sprintf("%s%t", k, vs)
		}

	}
	//fmt.Println(method+query)
	return strings.ToLower(encry.HmacHex(method+query, appSecret))
}
