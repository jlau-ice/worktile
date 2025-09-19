package main

import (
	"log"
	"worktile/worktile-query-server/routes"
	"worktile/worktile-query-server/server"
)

func main() {
	// 1. 创建并配置服务器实例
	e := server.NewServer()
	// 2. 注册所有路由
	routes.InitRoutes(e)
	// 3. 启动服务器
	log.Fatal(e.Start(":1323"))
}
