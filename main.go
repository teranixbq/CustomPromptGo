package main

import (
	"fmt"
	"promptgo/app/config"
	"promptgo/app/database"
	"promptgo/app/route"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	cfg := config.InitConfig()
	db := database.InitDBPostgres(cfg)
	database.AutoMigrate(db)

	// g.SetTrustedProxies([]string{})
	route.Run(g, db)

	port := strconv.Itoa(cfg.SERVERPORT)
	g.Run(fmt.Sprintf(":%s", port))
}
