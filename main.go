package main

import (
	"fmt"
	"promptgo/app/config"
	"promptgo/app/database"
	"promptgo/app/route"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	cfg := config.InitConfig()
	db := database.InitDBPostgres(cfg)
	database.AutoMigrate(db)

	g.Use(cors.Default())
	route.Run(g, db)

	g.Run(fmt.Sprintf(":%s", cfg.SERVERPORT))
}
