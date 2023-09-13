package main

import (
	"ssk/databases"
	"ssk/routers"
)

func main() {
	r := routers.GetRouter()

	// 设置静态文件目录
	r.Static("/json", "./json")

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run()

	databases.CloseMysql()
}
