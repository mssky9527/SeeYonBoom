package plugin

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"regexp"
)

func Info_Build(urlFlag string) {
	if urlFlag == "" {
		fmt.Println("请提供一个主机")
		return
	}

	// 拼接完整的URL
	fullURL := urlFlag + "/seeyon/main.do?method=showAbout"

	// 创建一个新的HTTP请求
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Printf("创建请求时出错: %v\n", err)
		return
	}

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

	// 使用正则表达式提取Build Id
	re := regexp.MustCompile(`B\d+\.\d+\.CTP\d+`)
	match := re.FindString(string(body))
	if match != "" {
		color.Green("Build Id: ")
		fmt.Println(match)
	} else {
		color.Red("Build Id 探测接口不存在!")
	}
}
