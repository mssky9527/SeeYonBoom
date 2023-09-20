package plugin

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"regexp"
)

func Info_Version(urlFlag string) {
	if urlFlag == "" {
		fmt.Println("请提供一个URL")
		return
	}

	// 拼接完整的URL
	fullURL := urlFlag + "/seeyon/index.jsp"

	// 创建一个新的HTTP请求，并指定HTTP版本为1.0
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Printf("创建请求时出错: %v\n", err)
		return
	}
	req.Proto = "HTTP/1.0"
	req.ProtoMajor = 1
	req.ProtoMinor = 0

	// 使用http客户端发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送请求时出错: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应内容时出错: %v\n", err)
		return
	}

	// 使用正则表达式提取版本信息
	re := regexp.MustCompile(`V=([^"]+)`) // 这个正则表达式匹配 "V=" 和其后面的非空白字符
	version := re.FindStringSubmatch(string(body))

	if len(version) > 1 {
		color.Green("Version:")
		fmt.Println(version[1])
	} else {
		color.Red("dbq, No Version.")
	}

}
