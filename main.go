package main

import (
	"log"
	"worktile/worktile-query-server/db"
	"worktile/worktile-query-server/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.InitConnection()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // 或者指定具体的前端域名，如 "http://localhost:3000"
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/api/users", handlers.SearchUsersHandler)
	e.GET("/api/workload", handlers.GetWorkloadHandler)
	log.Fatal(e.Start(":1323"))
}
