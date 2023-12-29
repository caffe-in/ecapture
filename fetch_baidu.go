package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// 创建一个http.Client实例，该实例将管理持久连接
	// 默认情况下，http.Client已经是配置好的，可以处理连接的持久性
	client := &http.Client{}

	// 无限循环，发送请求到百度
	for {
		// 创建http请求
		req, err := http.NewRequest("GET", "https://www.baidu.com", nil)
		if err != nil {
			fmt.Printf("Error creating request: %s\n", err)
			continue
		}

		// 发出请求
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error sending request: %s\n", err)
			continue
		}

		// 读取并丢弃响应体以免内存泄漏
		_, err = io.Copy(ioutil.Discard, resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %s\n", err)
		}

		// 关闭响应体
		resp.Body.Close()

		fmt.Println("Fetched baidu homepage")

		// 等待1秒钟
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
}
