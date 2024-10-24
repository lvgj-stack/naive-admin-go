package main

import (
	"time"

	"github.com/gin-gonic/gin"

	"naive-admin-go/config"
	"naive-admin-go/db"
	"naive-admin-go/router"
)

func main() {
	var Loc, _ = time.LoadLocation("Asia/Shanghai")
	time.Local = Loc
	app := gin.Default()
	config.Init()
	db.Init()
	router.Init(app)
	app.Run(":18085")
}
