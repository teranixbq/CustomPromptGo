package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func Run(c *gin.Engine,db *gorm.DB) {
	base := c.Group("/")
	
	RoutePrompt(base,db)
	RouteChat(base,db)
}