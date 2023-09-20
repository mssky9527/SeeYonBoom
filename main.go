package main

import (
	"SeeYonBoom/plugin"
	"flag"
)

func main() {
	urlFlag := flag.String("url", "", "请求的URL，例如：http://127.0.0.1")
	flag.Parse()

	plugin.DisplayURL(*urlFlag)
	plugin.Info_Version(*urlFlag)
	plugin.Info_Build(*urlFlag)
	plugin.Info_Cookies(*urlFlag)
	plugin.Info_readfile(*urlFlag)
	//plugin.CheckEndpointAndBruteForceUser(*urlFlag)
	plugin.BruteForcePasswords(*urlFlag)
	// ... 调用其他插件函数
}
