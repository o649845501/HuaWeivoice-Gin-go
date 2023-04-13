package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	router := gin.Default()
	// 定义路由
	router.POST("call", call)
	//打开日志文件
	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Fatalln("打开日志文件失败")
	}
	//设置日志输出
	log.SetOutput(logFile)
	// 启动服务器
	router.Run(":8080")
}
