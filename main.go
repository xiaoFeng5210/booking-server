package main

import (
	"booking/example"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	example.Run(r)

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
