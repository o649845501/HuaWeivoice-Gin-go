package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type AuthRequest struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

type AuthResponse struct {
	AccessToken string `json:"AccessToken"`
	ApplyType   string `json:"ApplyType"`
	CreateTime  string `json:"CreateTime"`
	Expires     string `json:"Expires"`
	Scope       string `json:"Scope"`
	AppKey      string `json:"AppKey"`
	UserID      string `json:"UserID"`
}

func handleAuthRequest() (AuthResponse, error) {
	var req AuthRequest
	req.AppKey = "2049230d-54b9-44e1-aa57-7423705a94b8"
	req.AppSecret = "2a5207165ddaa14bb8bf6eab9ec2c2ca"

	// 调用API获取access token
	var response AuthResponse
	token, err := getToken(req.AppKey, req.AppSecret)
	response = *token
	return response, err
}

func getToken(appKey string, appSecret string) (*AuthResponse, error) {
	// 构造请求body
	reqBody, err := json.Marshal(AuthRequest{AppKey: appKey, AppSecret: appSecret})
	if err != nil {
		return nil, err
	}

	// 构造请求头
	reqHeader := http.Header{}
	reqHeader.Set("Content-Type", "application/json")
	reqHeader.Set("X-Token-Expire", "600")

	// 构造请求
	req, err := http.NewRequest(http.MethodPost, "https://servicestage.aicccloud.com/apigovernance/api/oauth/tokenByAkSk", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header = reqHeader
	//println("请求结构是")
	// 发送请求
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	//解析响应
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//打印响应
	defer resp.Body.Close()

	// 解析响应
	var respBody AuthResponse
	err = json.NewDecoder(buf).Decode(&respBody)
	if err != nil {
		println("解析响应失败")
		return nil, err
	}

	return &respBody, nil
}
