package mioApiClient

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func (c *Client) DoRequest() []byte {
	// 创建一个http客户端
	client := &http.Client{}

	fmt.Println("client url: ", c.InterfaceInfo.Url)

	// 创建对应的请求
	newRequest, err := http.NewRequest(c.InterfaceInfo.Method, c.InterfaceInfo.Url, nil)
	if err != nil {
		fmt.Println("[request err] NewRequest ", err.Error())
		return nil
	}
	// 填入请求头相关的参数
	newRequest.Header.Set("accessKey", c.Header.AccessKey)
	newRequest.Header.Set("sign", c.Header.Sign)
	newRequest.Header.Set("timestamp", c.Header.Timestamp)
	newRequest.Header.Set("nonce", strconv.Itoa(rand.Int()))
	newRequest.Header.Set("body", c.Header.Body)
	//newRequest.Header.Set("body", "body")

	// 客户端执行请求
	response, err := client.Do(newRequest)
	if err != nil {
		fmt.Println("[do request err] client.Do ", err.Error())
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Body 关闭失败！", err.Error())
		}
	}(response.Body)

	content, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	log.Println("Status: ", response.Status)
	log.Println("Body: ", string(content))
	//var res utils.Response
	//err = json.Unmarshal(content, &res)
	//if err != nil {
	//	fmt.Println("[Unmarshal context err] ", err.Error())
	//	return nil
	//}
	//if res.Code != 0 {
	//	log.Fatalln("Request Body: ", c.Header.Body)
	//}
	return content
}
