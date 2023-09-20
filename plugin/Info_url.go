package plugin

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

// 初始化函数，设置随机数种子，确保随机效果
func init() {
	rand.Seed(time.Now().UnixNano())
}

// DisplayURL displays the provided URL in a random color.
func DisplayURL(url string) {
	if url == "" {
		fmt.Println("没有提供URL")
		return
	}

	colorFuncs := []func(format string, a ...interface{}) string{
		color.RedString,
		color.GreenString,
		color.YellowString,
		color.BlueString,
		color.MagentaString,
		color.CyanString,
		color.WhiteString,
	}

	// 随机选择一个颜色函数
	chosenColor := colorFuncs[rand.Intn(len(colorFuncs))]

	fmt.Println(chosenColor("输入的URL: %s", url))
}
