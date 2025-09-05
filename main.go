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
		AllowOrigins: []string{"http://127.0.0.1:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/api/users", handlers.SearchUsersHandler)
	e.GET("/api/workload/:userID", handlers.GetWorkloadHandler)
	log.Fatal(e.Start(":1323"))
}
