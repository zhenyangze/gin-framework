package helpers

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// EncodeURL 转义url或转义其他字符
func EncodeURL(_url string) string {
	return url.QueryEscape(_url)
}

// DecodeURL 解义url
func DecodeURL(_url string) (string, error) {
	return url.QueryUnescape(_url)
}

// GetUrlParam 获取url中的参数（非解码）
func GetUrlParam(_url string, _key string) (value string) {
	u, err := url.Parse(_url)
	values := u.Query()
	if err != nil {
		value = ""
	} else {
		value = values.Get(_key)
	}
	return
}

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func simpleDownload(url string, localFile string) string {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	out, err := os.Create(localFile)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return ""
	}
	return localFile
}

func ReferDownload(url string, site string, localFile string) string {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return ""
	}
	// 添加请求头
	req.Header.Add("referer", site)
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36 Edg/101.0.1210.53")
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err")
		return ""
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	out, err := os.Create(localFile)
	if err != nil {
		return ""
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return ""
	}
	return localFile
}
