package athttp

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func HttpRequest(method func() (*http.Request, error)) (*BaseResponse, error) {
	client := &http.Client{}
	req, err := method() // GenerateReqeust(bc,param)
	resp, err := client.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return nil, err
	}
	var br BaseResponse
	json.Unmarshal(body, &br)
	fmt.Println(string(body))
	//返回请求的结果
	return &br, nil
}

/**
请求接口成功后返回的信息
*/
type BaseResponse struct {
	Code int
	Msg  string
	Data interface{}
}

func Http0Request(method func() (*http.Request, error)) ([]byte, error) {
	client := &http.Client{}
	req, err := method() // GenerateReqeust(bc,param)
	resp, err := client.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return nil, err
	}
	//返回请求的结果
	return body, nil
}
