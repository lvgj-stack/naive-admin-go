package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"naive-admin-go/config"
	"naive-admin-go/db"
	"naive-admin-go/pkg/log"
	"naive-admin-go/router"
)

func main() {
	var Loc, _ = time.LoadLocation("Asia/Shanghai")
	time.Local = Loc
	log.InitLog()
	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // 允许所有来源
		// 或者您可以指定特定的来源
		// AllowOrigins: []string{"http://example.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	config.Init()
	db.Init()
	router.Init(app)
	app.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
