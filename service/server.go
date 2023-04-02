package service

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Request(method, url, data string) ([]byte, error) {
	var client *http.Client
	req, err := http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, err
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func RequestChat(msg, parentID string) string {
	url := "https://xinchain.ai/api/chat-process"
	method := "POST"
	bd := fmt.Sprintf(`{"prompt":"%s","options":{"parentMessageId":"%s"}}`, msg, parentID)
	fmt.Println("gpt请求体", bd)
	payload := strings.NewReader(bd)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("authority", "xinchain.ai")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cookie", "Hm_lvt_66386a83bdca11487201219600a58c56=1680432591; Hm_lpvt_66386a83bdca11487201219600a58c56=1680434762")
	req.Header.Add("origin", "https://xinchain.ai")
	req.Header.Add("referer", "https://xinchain.ai/")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"111\", \"Not(A:Brand\";v=\"8\", \"Chromium\";v=\"111\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}
