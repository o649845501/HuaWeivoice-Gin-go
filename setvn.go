package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func call(c *gin.Context) {
	println("____________________读取请求参数——————————————————————")
	voiceContent := c.PostForm("voiceContent")
	called := c.PostForm("called")
	tokens, _ := handleAuthRequest()
	appKey := "2049230d-54b9-44e1-aa57-7423705a94b8"
	accessToken := tokens.AccessToken
	//called以逗号为分割符切割存入数组
	calleds := strings.Split(called, ",")
	//循环数组，调用接口
	for _, called := range calleds {
		println("called是", called)
		println("voiceContent是", voiceContent)
		println("accessToken是", accessToken)
		println("appKey是", appKey)
		apiURL := "https://servicestage.aicccloud.com/apiaccess/rest/voiceNotification/v1/createVoiceNotification"
		timeout := time.Duration(10 * time.Second)

		reqBody := url.Values{}
		reqBody.Set("voiceContent", voiceContent)
		reqBody.Set("called", called)

		req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(reqBody.Encode()))
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("X-APP-Key", appKey)
		req.Header.Set("Authorization", "Bearer "+accessToken)

		client := &http.Client{Timeout: timeout}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp.Body)
		fmt.Println(string(buf.Bytes()))
		//等待1秒
		time.Sleep(1 * time.Second)
		log.Println(string(buf.Bytes()), time.Now().Format("2006-01-02 15:04:05"), "推送成功")
	}

}
